package artwork

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/core"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/utils"
)

type artistReader struct {
	cacheKey
	a            *artwork
	em           core.ExternalMetadata
	artist       model.Artist
	artistFolder string
	files        string
}

func newArtistReader(ctx context.Context, artwork *artwork, artID model.ArtworkID, em core.ExternalMetadata) (*artistReader, error) {
	ar, err := artwork.ds.Artist(ctx).Get(artID.ID)
	if err != nil {
		return nil, err
	}
	als, err := artwork.ds.Album(ctx).GetAll(model.QueryOptions{Filters: squirrel.Eq{"album_artist_id": artID.ID}})
	if err != nil {
		return nil, err
	}
	a := &artistReader{
		a:      artwork,
		em:     em,
		artist: *ar,
	}
	a.cacheKey.lastUpdate = ar.ExternalInfoUpdatedAt
	var files []string
	var paths []string
	for _, al := range als {
		files = append(files, al.ImageFiles)
		paths = append(paths, filepath.SplitList(al.Paths)...)
		if a.cacheKey.lastUpdate.Before(al.UpdatedAt) {
			a.cacheKey.lastUpdate = al.UpdatedAt
		}
	}
	a.files = strings.Join(files, string(filepath.ListSeparator))
	a.artistFolder = utils.LongestCommonPrefix(paths)
	if !strings.HasSuffix(a.artistFolder, string(filepath.Separator)) {
		a.artistFolder, _ = filepath.Split(a.artistFolder)
	}
	a.cacheKey.artID = artID
	return a, nil
}

func (a *artistReader) Key() string {
	hash := md5.Sum([]byte(conf.Server.Agents + conf.Server.Spotify.ID))
	return fmt.Sprintf(
		"%s.%x.%t",
		a.cacheKey.Key(),
		hash,
		conf.Server.EnableExternalServices,
	)
}

func (a *artistReader) LastUpdated() time.Time {
	return a.lastUpdate
}

func (a *artistReader) Reader(ctx context.Context) (io.ReadCloser, string, error) {
	return selectImageReader(ctx, a.artID,
		fromArtistFolder(ctx, a.artistFolder, "artist.*"),
		fromExternalFile(ctx, a.files, "artist.*"),
		fromArtistExternalSource(ctx, a.artist, a.em),
		fromArtistPlaceholder(),
	)
}

func fromArtistFolder(ctx context.Context, artistFolder string, pattern string) sourceFunc {
	return func() (io.ReadCloser, string, error) {
		fsys := os.DirFS(artistFolder)
		matches, err := fs.Glob(fsys, pattern)
		if err != nil {
			log.Warn(ctx, "Error matching artist image pattern", "pattern", pattern, "folder", artistFolder)
			return nil, "", err
		}
		if len(matches) == 0 {
			return nil, "", fmt.Errorf(`no matches for '%s' in '%s'`, pattern, artistFolder)
		}
		filePath := filepath.Join(artistFolder, matches[0])
		f, err := os.Open(filePath)
		if err != nil {
			log.Warn(ctx, "Could not open cover art file", "file", filePath, err)
			return nil, "", err
		}
		return f, filePath, err
	}
}
