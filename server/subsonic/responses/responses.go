package responses

import (
	"encoding/xml"
	"time"
)

type Subsonic struct {
	XMLName       xml.Name           `xml:"http://subsonic.org/restapi subsonic-response" json:"-"`
	Status        string             `xml:"status,attr"                                   json:"status"`
	Version       string             `xml:"version,attr"                                  json:"version"`
	Type          string             `xml:"type,attr"                                     json:"type"`
	ServerVersion string             `xml:"serverVersion,attr"                            json:"serverVersion"`
	Error         *Error             `xml:"error,omitempty"                               json:"error,omitempty"`
	License       *License           `xml:"license,omitempty"                             json:"license,omitempty"`
	MusicFolders  *MusicFolders      `xml:"musicFolders,omitempty"                        json:"musicFolders,omitempty"`
	Indexes       *Indexes           `xml:"indexes,omitempty"                             json:"indexes,omitempty"`
	Directory     *Directory         `xml:"directory,omitempty"                           json:"directory,omitempty"`
	User          *User              `xml:"user,omitempty"                                json:"user,omitempty"`
	Users         *Users             `xml:"users,omitempty"                               json:"users,omitempty"`
	AlbumList     *AlbumList         `xml:"albumList,omitempty"                           json:"albumList,omitempty"`
	AlbumList2    *AlbumList         `xml:"albumList2,omitempty"                          json:"albumList2,omitempty"`
	Playlists     *Playlists         `xml:"playlists,omitempty"                           json:"playlists,omitempty"`
	Playlist      *PlaylistWithSongs `xml:"playlist,omitempty"                            json:"playlist,omitempty"`
	SearchResult2 *SearchResult2     `xml:"searchResult2,omitempty"                       json:"searchResult2,omitempty"`
	SearchResult3 *SearchResult3     `xml:"searchResult3,omitempty"                       json:"searchResult3,omitempty"`
	Starred       *Starred           `xml:"starred,omitempty"                             json:"starred,omitempty"`
	Starred2      *Starred           `xml:"starred2,omitempty"                            json:"starred2,omitempty"`
	NowPlaying    *NowPlaying        `xml:"nowPlaying,omitempty"                          json:"nowPlaying,omitempty"`
	Song          *Child             `xml:"song,omitempty"                                json:"song,omitempty"`
	RandomSongs   *Songs             `xml:"randomSongs,omitempty"                         json:"randomSongs,omitempty"`
	SongsByGenre  *Songs             `xml:"songsByGenre,omitempty"                        json:"songsByGenre,omitempty"`
	Genres        *Genres            `xml:"genres,omitempty"                              json:"genres,omitempty"`

	// ID3
	Artist              *Indexes             `xml:"artists,omitempty"                     json:"artists,omitempty"`
	ArtistWithAlbumsID3 *ArtistWithAlbumsID3 `xml:"artist,omitempty"                      json:"artist,omitempty"`
	AlbumWithSongsID3   *AlbumWithSongsID3   `xml:"album,omitempty"                       json:"album,omitempty"`

	ArtistInfo    *ArtistInfo    `xml:"artistInfo,omitempty"                              json:"artistInfo,omitempty"`
	ArtistInfo2   *ArtistInfo2   `xml:"artistInfo2,omitempty"                             json:"artistInfo2,omitempty"`
	SimilarSongs  *SimilarSongs  `xml:"similarSongs,omitempty"                            json:"similarSongs,omitempty"`
	SimilarSongs2 *SimilarSongs2 `xml:"similarSongs2,omitempty"                           json:"similarSongs2,omitempty"`
	TopSongs      *TopSongs      `xml:"topSongs,omitempty"                                json:"topSongs,omitempty"`

	PlayQueue  *PlayQueue  `xml:"playQueue,omitempty"                                     json:"playQueue,omitempty"`
	Bookmarks  *Bookmarks  `xml:"bookmarks,omitempty"                                     json:"bookmarks,omitempty"`
	ScanStatus *ScanStatus `xml:"scanStatus,omitempty"                                    json:"scanStatus,omitempty"`
	Lyrics     *Lyrics     `xml:"lyrics,omitempty"                                        json:"lyrics,omitempty"`
}

type JsonWrapper struct {
	Subsonic Subsonic `json:"subsonic-response"`
}

type Error struct {
	Code    int    `xml:"code,attr"                      json:"code"`
	Message string `xml:"message,attr"                   json:"message"`
}

type License struct {
	Valid bool `xml:"valid,attr"                         json:"valid"`
}

type MusicFolder struct {
	Id   int32  `xml:"id,attr"                           json:"id"`
	Name string `xml:"name,attr"                         json:"name"`
}

type MusicFolders struct {
	Folders []MusicFolder `xml:"musicFolder"             json:"musicFolder,omitempty"`
}

type Artist struct {
	Id             string     `xml:"id,attr"                           json:"id"`
	Name           string     `xml:"name,attr"                         json:"name"`
	AlbumCount     int        `xml:"albumCount,attr,omitempty"         json:"albumCount,omitempty"`
	Starred        *time.Time `xml:"starred,attr,omitempty"            json:"starred,omitempty"`
	UserRating     int        `xml:"userRating,attr,omitempty"         json:"userRating,omitempty"`
	ArtistImageUrl string     `xml:"artistImageUrl,attr,omitempty"     json:"artistImageUrl,omitempty"`
	/*
		<xs:attribute name="averageRating" type="sub:AverageRating" use="optional"/>  <!-- Added in 1.13.0 -->
	*/
}

type Index struct {
	Name    string   `xml:"name,attr"                     json:"name"`
	Artists []Artist `xml:"artist"                        json:"artist"`
}

type Indexes struct {
	Index           []Index `xml:"index"                  json:"index,omitempty"`
	LastModified    int64   `xml:"lastModified,attr"      json:"lastModified"`
	IgnoredArticles string  `xml:"ignoredArticles,attr"   json:"ignoredArticles"`
}

type Child struct {
	Id                    string     `xml:"id,attr"                                 json:"id"`
	Parent                string     `xml:"parent,attr,omitempty"                   json:"parent,omitempty"`
	IsDir                 bool       `xml:"isDir,attr"                              json:"isDir"`
	Title                 string     `xml:"title,attr,omitempty"                    json:"title,omitempty"`
	Name                  string     `xml:"name,attr,omitempty"                     json:"name,omitempty"`
	Album                 string     `xml:"album,attr,omitempty"                    json:"album,omitempty"`
	Artist                string     `xml:"artist,attr,omitempty"                   json:"artist,omitempty"`
	Track                 int        `xml:"track,attr,omitempty"                    json:"track,omitempty"`
	Year                  int        `xml:"year,attr,omitempty"                     json:"year,omitempty"`
	Genre                 string     `xml:"genre,attr,omitempty"                    json:"genre,omitempty"`
	CoverArt              string     `xml:"coverArt,attr,omitempty"                 json:"coverArt,omitempty"`
	Size                  int64      `xml:"size,attr,omitempty"                     json:"size,omitempty"`
	ContentType           string     `xml:"contentType,attr,omitempty"              json:"contentType,omitempty"`
	Suffix                string     `xml:"suffix,attr,omitempty"                   json:"suffix,omitempty"`
	Starred               *time.Time `xml:"starred,attr,omitempty"                  json:"starred,omitempty"`
	TranscodedContentType string     `xml:"transcodedContentType,attr,omitempty"    json:"transcodedContentType,omitempty"`
	TranscodedSuffix      string     `xml:"transcodedSuffix,attr,omitempty"         json:"transcodedSuffix,omitempty"`
	Duration              int        `xml:"duration,attr,omitempty"                 json:"duration,omitempty"`
	BitRate               int        `xml:"bitRate,attr,omitempty"                  json:"bitRate,omitempty"`
	Path                  string     `xml:"path,attr,omitempty"                     json:"path,omitempty"`
	PlayCount             int64      `xml:"playCount,attr,omitempty"                json:"playCount,omitempty"`
	Played                *time.Time `xml:"played,attr,omitempty"                   json:"played,omitempty"`
	DiscNumber            int        `xml:"discNumber,attr,omitempty"               json:"discNumber,omitempty"`
	Created               *time.Time `xml:"created,attr,omitempty"                  json:"created,omitempty"`
	AlbumId               string     `xml:"albumId,attr,omitempty"                  json:"albumId,omitempty"`
	ArtistId              string     `xml:"artistId,attr,omitempty"                 json:"artistId,omitempty"`
	Type                  string     `xml:"type,attr,omitempty"                     json:"type,omitempty"`
	UserRating            int        `xml:"userRating,attr,omitempty"               json:"userRating,omitempty"`
	SongCount             int        `xml:"songCount,attr,omitempty"                json:"songCount,omitempty"`
	IsVideo               bool       `xml:"isVideo,attr"                            json:"isVideo"`
	BookmarkPosition      int64      `xml:"bookmarkPosition,attr,omitempty"         json:"bookmarkPosition,omitempty"`
	/*
	   <xs:attribute name="averageRating" type="sub:AverageRating" use="optional"/>  <!-- Added in 1.6.0 -->
	*/
}

type Songs struct {
	Songs []Child `xml:"song"                              json:"song,omitempty"`
}

type Directory struct {
	Child      []Child    `xml:"child"                              json:"child,omitempty"`
	Id         string     `xml:"id,attr"                            json:"id"`
	Name       string     `xml:"name,attr"                          json:"name"`
	Parent     string     `xml:"parent,attr,omitempty"              json:"parent,omitempty"`
	Starred    *time.Time `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	PlayCount  int64      `xml:"playCount,attr,omitempty"           json:"playCount,omitempty"`
	Played     *time.Time `xml:"played,attr,omitempty"              json:"played,omitempty"`
	UserRating int        `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`

	// ID3
	Artist     string     `xml:"artist,attr,omitempty"              json:"artist,omitempty"`
	ArtistId   string     `xml:"artistId,attr,omitempty"            json:"artistId,omitempty"`
	CoverArt   string     `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	SongCount  int        `xml:"songCount,attr,omitempty"           json:"songCount,omitempty"`
	AlbumCount int        `xml:"albumCount,attr,omitempty"          json:"albumCount,omitempty"`
	Duration   int        `xml:"duration,attr,omitempty"            json:"duration,omitempty"`
	Created    *time.Time `xml:"created,attr,omitempty"             json:"created,omitempty"`
	Year       int        `xml:"year,attr,omitempty"                json:"year,omitempty"`
	Genre      string     `xml:"genre,attr,omitempty"               json:"genre,omitempty"`

	/*
	   <xs:attribute name="averageRating" type="sub:AverageRating" use="optional"/>  <!-- Added in 1.13.0 -->
	*/
}

type ArtistID3 struct {
	Id             string     `xml:"id,attr"                            json:"id"`
	Name           string     `xml:"name,attr"                          json:"name"`
	CoverArt       string     `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	AlbumCount     int        `xml:"albumCount,attr,omitempty"          json:"albumCount,omitempty"`
	Starred        *time.Time `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	UserRating     int        `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`
	ArtistImageUrl string     `xml:"artistImageUrl,attr,omitempty"      json:"artistImageUrl,omitempty"`
}

type AlbumID3 struct {
	Id         string     `xml:"id,attr"                            json:"id"`
	Name       string     `xml:"name,attr"                          json:"name"`
	Artist     string     `xml:"artist,attr,omitempty"              json:"artist,omitempty"`
	ArtistId   string     `xml:"artistId,attr,omitempty"            json:"artistId,omitempty"`
	CoverArt   string     `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	SongCount  int        `xml:"songCount,attr,omitempty"           json:"songCount,omitempty"`
	Duration   int        `xml:"duration,attr,omitempty"            json:"duration,omitempty"`
	PlayCount  int64      `xml:"playCount,attr,omitempty"           json:"playCount,omitempty"`
	Played     *time.Time `xml:"played,attr,omitempty"              json:"played,omitempty"`
	Created    *time.Time `xml:"created,attr,omitempty"             json:"created,omitempty"`
	Starred    *time.Time `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	UserRating int        `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`
	Year       int        `xml:"year,attr,omitempty"                json:"year,omitempty"`
	Genre      string     `xml:"genre,attr,omitempty"               json:"genre,omitempty"`
}

type ArtistWithAlbumsID3 struct {
	ArtistID3
	Album []Child `xml:"album"                              json:"album,omitempty"`
}

type AlbumWithSongsID3 struct {
	AlbumID3
	Song []Child `xml:"song"                               json:"song,omitempty"`
}

type AlbumList struct {
	Album []Child `xml:"album"                                      json:"album,omitempty"`
}

type Playlist struct {
	Id        string    `xml:"id,attr"                       json:"id"`
	Name      string    `xml:"name,attr"                     json:"name"`
	Comment   string    `xml:"comment,attr,omitempty"        json:"comment,omitempty"`
	SongCount int       `xml:"songCount,attr"                json:"songCount"`
	Duration  int       `xml:"duration,attr"                 json:"duration"`
	Public    bool      `xml:"public,attr"                   json:"public"`
	Owner     string    `xml:"owner,attr,omitempty"          json:"owner,omitempty"`
	Created   time.Time `xml:"created,attr"                  json:"created"`
	Changed   time.Time `xml:"changed,attr"                  json:"changed"`
	/*
		<xs:sequence>
		    <xs:element name="allowedUser" type="xs:string" minOccurs="0" maxOccurs="unbounded"/> <!--Added in 1.8.0-->
		</xs:sequence>
		<xs:attribute name="coverArt" type="xs:string" use="optional"/>  <!--Added in 1.11.0-->
	*/
}

type Playlists struct {
	Playlist []Playlist `xml:"playlist"                           json:"playlist,omitempty"`
}

type PlaylistWithSongs struct {
	Playlist
	Entry []Child `xml:"entry"                                    json:"entry,omitempty"`
}

type SearchResult2 struct {
	Artist []Artist `xml:"artist"                                 json:"artist,omitempty"`
	Album  []Child  `xml:"album"                                  json:"album,omitempty"`
	Song   []Child  `xml:"song"                                   json:"song,omitempty"`
}

type SearchResult3 struct {
	Artist []ArtistID3 `xml:"artist"                                 json:"artist,omitempty"`
	Album  []Child     `xml:"album"                                  json:"album,omitempty"`
	Song   []Child     `xml:"song"                                   json:"song,omitempty"`
}

type Starred struct {
	Artist []Artist `xml:"artist"                                 json:"artist,omitempty"`
	Album  []Child  `xml:"album"                                  json:"album,omitempty"`
	Song   []Child  `xml:"song"                                   json:"song,omitempty"`
}

type NowPlayingEntry struct {
	Child
	UserName   string `xml:"username,attr"                        json:"username"`
	MinutesAgo int    `xml:"minutesAgo,attr"                      json:"minutesAgo"`
	PlayerId   int    `xml:"playerId,attr"                        json:"playerId"`
	PlayerName string `xml:"playerName,attr"                      json:"playerName,omitempty"`
}

type NowPlaying struct {
	Entry []NowPlayingEntry `xml:"entry"                          json:"entry,omitempty"`
}

type User struct {
	Username            string `xml:"username,attr"               json:"username"`
	Email               string `xml:"email,attr,omitempty"        json:"email,omitempty"`
	ScrobblingEnabled   bool   `xml:"scrobblingEnabled,attr"      json:"scrobblingEnabled"`
	MaxBitRate          int    `xml:"maxBitRate,attr,omitempty"   json:"maxBitRate,omitempty"`
	AdminRole           bool   `xml:"adminRole,attr"              json:"adminRole"`
	SettingsRole        bool   `xml:"settingsRole,attr"           json:"settingsRole"`
	DownloadRole        bool   `xml:"downloadRole,attr"           json:"downloadRole"`
	UploadRole          bool   `xml:"uploadRole,attr"             json:"uploadRole"`
	PlaylistRole        bool   `xml:"playlistRole,attr"           json:"playlistRole"`
	CoverArtRole        bool   `xml:"coverArtRole,attr"           json:"coverArtRole"`
	CommentRole         bool   `xml:"commentRole,attr"            json:"commentRole"`
	PodcastRole         bool   `xml:"podcastRole,attr"            json:"podcastRole"`
	StreamRole          bool   `xml:"streamRole,attr"             json:"streamRole"`
	JukeboxRole         bool   `xml:"jukeboxRole,attr"            json:"jukeboxRole"`
	ShareRole           bool   `xml:"shareRole,attr"              json:"shareRole"`
	VideoConversionRole bool   `xml:"videoConversionRole,attr"    json:"videoConversionRole"`
	Folder              []int  `xml:"folder,omitempty"            json:"folder,omitempty"`
}

type Users struct {
	User []User `xml:"user"  json:"user"`
}

type Genre struct {
	Name       string `xml:",chardata"                      json:"value,omitempty"`
	SongCount  int    `xml:"songCount,attr"             json:"songCount"`
	AlbumCount int    `xml:"albumCount,attr"            json:"albumCount"`
}

type Genres struct {
	Genre []Genre `xml:"genre,omitempty"                      json:"genre,omitempty"`
}

type ArtistInfoBase struct {
	Biography      string `xml:"biography,omitempty"          json:"biography,omitempty"`
	MusicBrainzID  string `xml:"musicBrainzId,omitempty"      json:"musicBrainzId,omitempty"`
	LastFmUrl      string `xml:"lastFmUrl,omitempty"          json:"lastFmUrl,omitempty"`
	SmallImageUrl  string `xml:"smallImageUrl,omitempty"      json:"smallImageUrl,omitempty"`
	MediumImageUrl string `xml:"mediumImageUrl,omitempty"     json:"mediumImageUrl,omitempty"`
	LargeImageUrl  string `xml:"largeImageUrl,omitempty"      json:"largeImageUrl,omitempty"`
}

type ArtistInfo struct {
	ArtistInfoBase
	SimilarArtist []Artist `xml:"similarArtist,omitempty"    json:"similarArtist,omitempty"`
}

type ArtistInfo2 struct {
	ArtistInfoBase
	SimilarArtist []ArtistID3 `xml:"similarArtist,omitempty"    json:"similarArtist,omitempty"`
}

type SimilarSongs struct {
	Song []Child `xml:"song,omitempty"         json:"song,omitempty"`
}

type SimilarSongs2 struct {
	Song []Child `xml:"song,omitempty"         json:"song,omitempty"`
}

type TopSongs struct {
	Song []Child `xml:"song,omitempty"         json:"song,omitempty"`
}

type PlayQueue struct {
	Entry     []Child    `xml:"entry,omitempty"         json:"entry,omitempty"`
	Current   string     `xml:"current,attr,omitempty"  json:"current,omitempty"`
	Position  int64      `xml:"position,attr,omitempty" json:"position,omitempty"`
	Username  string     `xml:"username,attr"           json:"username"`
	Changed   *time.Time `xml:"changed,attr,omitempty"  json:"changed,omitempty"`
	ChangedBy string     `xml:"changedBy,attr"          json:"changedBy"`
}

type Bookmark struct {
	Entry    Child     `xml:"entry,omitempty"         json:"entry,omitempty"`
	Position int64     `xml:"position,attr,omitempty" json:"position,omitempty"`
	Username string    `xml:"username,attr"           json:"username"`
	Comment  string    `xml:"comment,attr"            json:"comment"`
	Created  time.Time `xml:"created,attr"            json:"created"`
	Changed  time.Time `xml:"changed,attr"            json:"changed"`
}

type Bookmarks struct {
	Bookmark []Bookmark `xml:"bookmark,omitempty"    json:"bookmark,omitempty"`
}

type ScanStatus struct {
	Scanning    bool       `xml:"scanning,attr"            json:"scanning"`
	Count       int64      `xml:"count,attr"               json:"count"`
	FolderCount int64      `xml:"folderCount,attr"         json:"folderCount"`
	LastScan    *time.Time `xml:"lastScan,attr,omitempty"  json:"lastScan,omitempty"`
}

type Lyrics struct {
	Artist string `xml:"artist,omitempty,attr"  json:"artist,omitempty"`
	Title  string `xml:"title,omitempty,attr"   json:"title,omitempty"`
	Value  string `xml:",chardata"              json:"value"`
}
