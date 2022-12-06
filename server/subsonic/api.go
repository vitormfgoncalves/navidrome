package subsonic

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/navidrome/navidrome/consts"
	"github.com/navidrome/navidrome/core"
	"github.com/navidrome/navidrome/core/scrobbler"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/scanner"
	"github.com/navidrome/navidrome/server/events"
	"github.com/navidrome/navidrome/server/subsonic/responses"
	"github.com/navidrome/navidrome/utils"
)

const Version = "1.16.1"

type handler = func(*http.Request) (*responses.Subsonic, error)
type handlerRaw = func(http.ResponseWriter, *http.Request) (*responses.Subsonic, error)

type Router struct {
	http.Handler
	ds               model.DataStore
	artwork          core.Artwork
	streamer         core.MediaStreamer
	archiver         core.Archiver
	players          core.Players
	externalMetadata core.ExternalMetadata
	playlists        core.Playlists
	scanner          scanner.Scanner
	broker           events.Broker
	scrobbler        scrobbler.PlayTracker
}

func New(ds model.DataStore, artwork core.Artwork, streamer core.MediaStreamer, archiver core.Archiver,
	players core.Players, externalMetadata core.ExternalMetadata, scanner scanner.Scanner, broker events.Broker,
	playlists core.Playlists, scrobbler scrobbler.PlayTracker) *Router {
	r := &Router{
		ds:               ds,
		artwork:          artwork,
		streamer:         streamer,
		archiver:         archiver,
		players:          players,
		externalMetadata: externalMetadata,
		playlists:        playlists,
		scanner:          scanner,
		broker:           broker,
		scrobbler:        scrobbler,
	}
	r.Handler = r.routes()
	return r
}

func (api *Router) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(postFormToQueryParams)
	r.Use(checkRequiredParameters)
	r.Use(authenticate(api.ds))
	// TODO Validate version

	// Subsonic endpoints, grouped by controller
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "ping", api.Ping)
		h(r, "getLicense", api.GetLicense)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "getMusicFolders", api.GetMusicFolders)
		h(r, "getIndexes", api.GetIndexes)
		h(r, "getArtists", api.GetArtists)
		h(r, "getGenres", api.GetGenres)
		h(r, "getMusicDirectory", api.GetMusicDirectory)
		h(r, "getArtist", api.GetArtist)
		h(r, "getAlbum", api.GetAlbum)
		h(r, "getSong", api.GetSong)
		h(r, "getArtistInfo", api.GetArtistInfo)
		h(r, "getArtistInfo2", api.GetArtistInfo2)
		h(r, "getTopSongs", api.GetTopSongs)
		h(r, "getSimilarSongs", api.GetSimilarSongs)
		h(r, "getSimilarSongs2", api.GetSimilarSongs2)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		hr(r, "getAlbumList", api.GetAlbumList)
		hr(r, "getAlbumList2", api.GetAlbumList2)
		h(r, "getStarred", api.GetStarred)
		h(r, "getStarred2", api.GetStarred2)
		h(r, "getNowPlaying", api.GetNowPlaying)
		h(r, "getRandomSongs", api.GetRandomSongs)
		h(r, "getSongsByGenre", api.GetSongsByGenre)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "setRating", api.SetRating)
		h(r, "star", api.Star)
		h(r, "unstar", api.Unstar)
		h(r, "scrobble", api.Scrobble)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "getPlaylists", api.GetPlaylists)
		h(r, "getPlaylist", api.GetPlaylist)
		h(r, "createPlaylist", api.CreatePlaylist)
		h(r, "deletePlaylist", api.DeletePlaylist)
		h(r, "updatePlaylist", api.UpdatePlaylist)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "getBookmarks", api.GetBookmarks)
		h(r, "createBookmark", api.CreateBookmark)
		h(r, "deleteBookmark", api.DeleteBookmark)
		h(r, "getPlayQueue", api.GetPlayQueue)
		h(r, "savePlayQueue", api.SavePlayQueue)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		h(r, "search2", api.Search2)
		h(r, "search3", api.Search3)
	})
	r.Group(func(r chi.Router) {
		h(r, "getUser", api.GetUser)
		h(r, "getUsers", api.GetUsers)
	})
	r.Group(func(r chi.Router) {
		h(r, "getScanStatus", api.GetScanStatus)
		h(r, "startScan", api.StartScan)
	})
	r.Group(func(r chi.Router) {
		// configure request throttling
		maxRequests := utils.MaxInt(2, runtime.NumCPU())
		r.Use(middleware.ThrottleBacklog(maxRequests, consts.RequestThrottleBacklogLimit, consts.RequestThrottleBacklogTimeout))
		hr(r, "getAvatar", api.GetAvatar)
		hr(r, "getCoverArt", api.GetCoverArt)
		h(r, "getLyrics", api.GetLyrics)
	})
	r.Group(func(r chi.Router) {
		r.Use(getPlayer(api.players))
		hr(r, "stream", api.Stream)
		hr(r, "download", api.Download)
	})

	// Not Implemented (yet?)
	h501(r, "jukeboxControl")
	h501(r, "getAlbumInfo", "getAlbumInfo2")
	h501(r, "getShares", "createShare", "updateShare", "deleteShare")
	h501(r, "getPodcasts", "getNewestPodcasts", "refreshPodcasts", "createPodcastChannel", "deletePodcastChannel",
		"deletePodcastEpisode", "downloadPodcastEpisode")
	h501(r, "getInternetRadioStations", "createInternetRadioStation", "updateInternetRadioStation",
		"deleteInternetRadioStation")
	h501(r, "createUser", "updateUser", "deleteUser", "changePassword")

	// Deprecated/Won't implement/Out of scope endpoints
	h410(r, "search")
	h410(r, "getChatMessages", "addChatMessage")
	h410(r, "getVideos", "getVideoInfo", "getCaptions", "hls")
	return r
}

// Add a Subsonic handler
func h(r chi.Router, path string, f handler) {
	hr(r, path, func(_ http.ResponseWriter, r *http.Request) (*responses.Subsonic, error) {
		return f(r)
	})
}

// Add a Subsonic handler that requires a http.ResponseWriter (ex: stream, getCoverArt...)
func hr(r chi.Router, path string, f handlerRaw) {
	handle := func(w http.ResponseWriter, r *http.Request) {
		res, err := f(w, r)
		if err != nil {
			// If it is not a Subsonic error, convert it to an ErrorGeneric
			var subErr subError
			if !errors.As(err, &subErr) {
				if errors.Is(err, model.ErrNotFound) {
					err = newError(responses.ErrorDataNotFound, "data not found")
				} else {
					err = newError(responses.ErrorGeneric, "Internal Error")
				}
			}
			sendError(w, r, err)
			return
		}
		if r.Context().Err() != nil {
			if log.CurrentLevel() >= log.LevelDebug {
				log.Warn(r.Context(), "Request was interrupted", "path", path, r.Context().Err())
			}
			return
		}
		if res != nil {
			sendResponse(w, r, res)
		}
	}
	addHandler(r, path, handle)
}

// Add a handler that returns 501 - Not implemented. Used to signal that an endpoint is not implemented yet
func h501(r chi.Router, paths ...string) {
	for _, path := range paths {
		handle := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cache-Control", "no-cache")
			w.WriteHeader(501)
			_, _ = w.Write([]byte("This endpoint is not implemented, but may be in future releases"))
		}
		addHandler(r, path, handle)
	}
}

// Add a handler that returns 410 - Gone. Used to signal that an endpoint will not be implemented
func h410(r chi.Router, paths ...string) {
	for _, path := range paths {
		handle := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(410)
			_, _ = w.Write([]byte("This endpoint will not be implemented"))
		}
		addHandler(r, path, handle)
	}
}

func addHandler(r chi.Router, path string, handle func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc("/"+path, handle)
	r.HandleFunc("/"+path+".view", handle)
}

func sendError(w http.ResponseWriter, r *http.Request, err error) {
	response := newResponse()
	code := responses.ErrorGeneric
	var subErr subError
	if errors.As(err, &subErr) {
		code = subErr.code
	}
	response.Status = "failed"
	response.Error = &responses.Error{Code: code, Message: err.Error()}

	sendResponse(w, r, response)
}

func sendResponse(w http.ResponseWriter, r *http.Request, payload *responses.Subsonic) {
	f := utils.ParamString(r, "f")
	var response []byte
	switch f {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		wrapper := &responses.JsonWrapper{Subsonic: *payload}
		response, _ = json.Marshal(wrapper)
	case "jsonp":
		w.Header().Set("Content-Type", "application/javascript")
		callback := utils.ParamString(r, "callback")
		wrapper := &responses.JsonWrapper{Subsonic: *payload}
		data, _ := json.Marshal(wrapper)
		response = []byte(fmt.Sprintf("%s(%s)", callback, data))
	default:
		w.Header().Set("Content-Type", "application/xml")
		response, _ = xml.Marshal(payload)
	}
	if payload.Status == "ok" {
		if log.CurrentLevel() >= log.LevelTrace {
			log.Debug(r.Context(), "API: Successful response", "status", "OK", "body", string(response))
		} else {
			log.Debug(r.Context(), "API: Successful response", "status", "OK")
		}
	} else {
		log.Warn(r.Context(), "API: Failed response", "error", payload.Error.Code, "message", payload.Error.Message)
	}
	if _, err := w.Write(response); err != nil {
		log.Error(r, "Error sending response to client", "payload", string(response), err)
	}
}
