package scrobbler

import (
	"context"
	"sort"
	"time"

	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/consts"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/model/request"
	"github.com/navidrome/navidrome/server/events"
	"github.com/navidrome/navidrome/utils/singleton"
)

const nowPlayingExpire = 60 * time.Minute

type NowPlayingInfo struct {
	TrackID    string
	Start      time.Time
	Username   string
	PlayerId   string
	PlayerName string
}

type Submission struct {
	TrackID   string
	Timestamp time.Time
}

type PlayTracker interface {
	NowPlaying(ctx context.Context, playerId string, playerName string, trackId string) error
	GetNowPlaying(ctx context.Context) ([]NowPlayingInfo, error)
	Submit(ctx context.Context, submissions []Submission) error
}

type playTracker struct {
	ds         model.DataStore
	broker     events.Broker
	playMap    *ttlcache.Cache
	scrobblers map[string]Scrobbler
}

func GetPlayTracker(ds model.DataStore, broker events.Broker) PlayTracker {
	return singleton.GetInstance(func() *playTracker {
		m := ttlcache.NewCache()
		m.SkipTTLExtensionOnHit(true)
		_ = m.SetTTL(nowPlayingExpire)
		p := &playTracker{ds: ds, playMap: m, broker: broker}
		p.scrobblers = make(map[string]Scrobbler)
		for name, constructor := range constructors {
			s := constructor(ds)
			if conf.Server.DevEnableBufferedScrobble {
				s = newBufferedScrobbler(ds, s, name)
			}
			p.scrobblers[name] = s
		}
		return p
	})
}

func (p *playTracker) NowPlaying(ctx context.Context, playerId string, playerName string, trackId string) error {
	user, _ := request.UserFrom(ctx)
	info := NowPlayingInfo{
		TrackID:    trackId,
		Start:      time.Now(),
		Username:   user.UserName,
		PlayerId:   playerId,
		PlayerName: playerName,
	}
	_ = p.playMap.Set(playerId, info)
	player, _ := request.PlayerFrom(ctx)
	if player.ScrobbleEnabled {
		p.dispatchNowPlaying(ctx, user.ID, trackId)
	}
	return nil
}

func (p *playTracker) dispatchNowPlaying(ctx context.Context, userId string, trackId string) {
	t, err := p.ds.MediaFile(ctx).Get(trackId)
	if err != nil {
		log.Error(ctx, "Error retrieving mediaFile", "id", trackId, err)
		return
	}
	if t.Artist == consts.UnknownArtist {
		log.Debug(ctx, "Ignoring external NowPlaying update for track with unknown artist", "track", t.Title, "artist", t.Artist)
		return
	}
	// TODO Parallelize
	for name, s := range p.scrobblers {
		if !s.IsAuthorized(ctx, userId) {
			continue
		}
		log.Debug(ctx, "Sending NowPlaying update", "scrobbler", name, "track", t.Title, "artist", t.Artist)
		err := s.NowPlaying(ctx, userId, t)
		if err != nil {
			log.Error(ctx, "Error sending NowPlayingInfo", "scrobbler", name, "track", t.Title, "artist", t.Artist, err)
			continue
		}
	}
}

func (p *playTracker) GetNowPlaying(ctx context.Context) ([]NowPlayingInfo, error) {
	var res []NowPlayingInfo
	for _, playerId := range p.playMap.GetKeys() {
		value, err := p.playMap.Get(playerId)
		if err != nil {
			continue
		}
		info := value.(NowPlayingInfo)
		res = append(res, info)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Start.After(res[j].Start)
	})
	return res, nil
}

func (p *playTracker) Submit(ctx context.Context, submissions []Submission) error {
	username, _ := request.UsernameFrom(ctx)
	player, _ := request.PlayerFrom(ctx)
	if !player.ScrobbleEnabled {
		log.Debug(ctx, "External scrobbling disabled for this player", "player", player.Name, "ip", player.IPAddress, "user", username)
	}
	event := &events.RefreshResource{}
	success := 0

	for _, s := range submissions {
		mf, err := p.ds.MediaFile(ctx).Get(s.TrackID)
		if err != nil {
			log.Error(ctx, "Cannot find track for scrobbling", "id", s.TrackID, "user", username, err)
			continue
		}
		err = p.incPlay(ctx, mf, s.Timestamp)
		if err != nil {
			log.Error(ctx, "Error updating play counts", "id", mf.ID, "track", mf.Title, "user", username, err)
		} else {
			success++
			event.With("song", mf.ID).With("album", mf.AlbumID).With("artist", mf.AlbumArtistID)
			log.Info(ctx, "Scrobbled", "title", mf.Title, "artist", mf.Artist, "user", username, "timestamp", s.Timestamp)
			if player.ScrobbleEnabled {
				p.dispatchScrobble(ctx, mf, s.Timestamp)
			}
		}
	}

	if success > 0 {
		p.broker.SendMessage(ctx, event)
	}
	return nil
}

func (p *playTracker) incPlay(ctx context.Context, track *model.MediaFile, timestamp time.Time) error {
	return p.ds.WithTx(func(tx model.DataStore) error {
		err := p.ds.MediaFile(ctx).IncPlayCount(track.ID, timestamp)
		if err != nil {
			return err
		}
		err = p.ds.Album(ctx).IncPlayCount(track.AlbumID, timestamp)
		if err != nil {
			return err
		}
		err = p.ds.Artist(ctx).IncPlayCount(track.ArtistID, timestamp)
		return err
	})
}

func (p *playTracker) dispatchScrobble(ctx context.Context, t *model.MediaFile, playTime time.Time) {
	if t.Artist == consts.UnknownArtist {
		log.Debug(ctx, "Ignoring external Scrobble for track with unknown artist", "track", t.Title, "artist", t.Artist)
		return
	}
	u, _ := request.UserFrom(ctx)
	scrobble := Scrobble{MediaFile: *t, TimeStamp: playTime}
	for name, s := range p.scrobblers {
		if !s.IsAuthorized(ctx, u.ID) {
			continue
		}
		if conf.Server.DevEnableBufferedScrobble {
			log.Debug(ctx, "Buffering Scrobble", "scrobbler", name, "track", t.Title, "artist", t.Artist)
		} else {
			log.Debug(ctx, "Sending Scrobble", "scrobbler", name, "track", t.Title, "artist", t.Artist)
		}
		err := s.Scrobble(ctx, u.ID, scrobble)
		if err != nil {
			log.Error(ctx, "Error sending Scrobble", "scrobbler", name, "track", t.Title, "artist", t.Artist, err)
			continue
		}
	}
}

var constructors map[string]Constructor

func Register(name string, init Constructor) {
	if constructors == nil {
		constructors = make(map[string]Constructor)
	}
	constructors[name] = init
}
