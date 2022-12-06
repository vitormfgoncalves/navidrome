//go:build unix

// TODO Fix snapshot tests in Windows
// Response Snapshot tests. Only run in Linux and macOS, as they fail in Windows
// Probably because of EOL char differences
package responses_test

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/navidrome/navidrome/consts"
	. "github.com/navidrome/navidrome/server/subsonic/responses"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Responses", func() {
	var response *Subsonic
	BeforeEach(func() {
		response = &Subsonic{Status: "ok", Version: "1.8.0", Type: consts.AppName, ServerVersion: "v0.0.0"}
	})

	Describe("EmptyResponse", func() {
		It("should match .XML", func() {
			Expect(xml.Marshal(response)).To(MatchSnapshot())
		})
		It("should match .JSON", func() {
			Expect(json.Marshal(response)).To(MatchSnapshot())
		})
	})

	Describe("License", func() {
		BeforeEach(func() {
			response.License = &License{Valid: true}
		})
		It("should match .XML", func() {
			Expect(xml.Marshal(response)).To(MatchSnapshot())
		})
		It("should match .JSON", func() {
			Expect(json.Marshal(response)).To(MatchSnapshot())
		})
	})

	Describe("MusicFolders", func() {
		BeforeEach(func() {
			response.MusicFolders = &MusicFolders{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				folders := make([]MusicFolder, 2)
				folders[0] = MusicFolder{Id: 111, Name: "aaa"}
				folders[1] = MusicFolder{Id: 222, Name: "bbb"}
				response.MusicFolders.Folders = folders
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Indexes", func() {
		BeforeEach(func() {
			response.Indexes = &Indexes{LastModified: 1, IgnoredArticles: "A"}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				artists := make([]Artist, 1)
				t := time.Date(2016, 03, 2, 20, 30, 0, 0, time.UTC)
				artists[0] = Artist{
					Id:             "111",
					Name:           "aaa",
					Starred:        &t,
					UserRating:     3,
					AlbumCount:     2,
					ArtistImageUrl: "https://lastfm.freetls.fastly.net/i/u/300x300/2a96cbd8b46e442fc41c2b86b821562f.png",
				}
				index := make([]Index, 1)
				index[0] = Index{Name: "A", Artists: artists}
				response.Indexes.Index = index
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Child", func() {
		Context("with data", func() {
			BeforeEach(func() {
				response.Directory = &Directory{Id: "1", Name: "N"}
				child := make([]Child, 1)
				t := time.Date(2016, 03, 2, 20, 30, 0, 0, time.UTC)
				child[0] = Child{
					Id: "1", IsDir: true, Title: "title", Album: "album", Artist: "artist", Track: 1,
					Year: 1985, Genre: "Rock", CoverArt: "1", Size: 8421341, ContentType: "audio/flac",
					Suffix: "flac", TranscodedContentType: "audio/mpeg", TranscodedSuffix: "mp3",
					Duration: 146, BitRate: 320, Starred: &t,
				}
				response.Directory.Child = child
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Directory", func() {
		BeforeEach(func() {
			response.Directory = &Directory{Id: "1", Name: "N"}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.Directory.Child = child
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("AlbumList", func() {
		BeforeEach(func() {
			response.AlbumList = &AlbumList{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.AlbumList.Album = child
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("User", func() {
		BeforeEach(func() {
			response.User = &User{Username: "deluan"}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				response.User.Email = "navidrome@deluan.com"
				response.User.Folder = []int{1}
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Users", func() {
		BeforeEach(func() {
			u := User{Username: "deluan"}
			response.Users = &Users{User: []User{u}}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				u := User{Username: "deluan"}
				u.Email = "navidrome@deluan.com"
				u.AdminRole = true
				u.Folder = []int{1}
				response.Users = &Users{User: []User{u}}
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Playlists", func() {
		BeforeEach(func() {
			response.Playlists = &Playlists{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			timestamp, _ := time.Parse(time.RFC3339, "2020-04-11T16:43:00Z04:00")
			BeforeEach(func() {
				pls := make([]Playlist, 2)
				pls[0] = Playlist{
					Id:        "111",
					Name:      "aaa",
					Comment:   "comment",
					SongCount: 2,
					Duration:  120,
					Public:    true,
					Owner:     "admin",
					Created:   timestamp,
					Changed:   timestamp,
				}
				pls[1] = Playlist{Id: "222", Name: "bbb"}
				response.Playlists.Playlist = pls
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Genres", func() {
		BeforeEach(func() {
			response.Genres = &Genres{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				genres := make([]Genre, 3)
				genres[0] = Genre{SongCount: 1000, AlbumCount: 100, Name: "Rock"}
				genres[1] = Genre{SongCount: 500, AlbumCount: 50, Name: "Reggae"}
				genres[2] = Genre{SongCount: 0, AlbumCount: 0, Name: "Pop"}
				response.Genres.Genre = genres
			})

			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("ArtistInfo", func() {
		BeforeEach(func() {
			response.ArtistInfo = &ArtistInfo{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				response.ArtistInfo.Biography = `Black Sabbath is an English <a target='_blank' href="http://www.last.fm/tag/heavy%20metal" class="bbcode_tag" rel="tag">heavy metal</a> band`
				response.ArtistInfo.MusicBrainzID = "5182c1d9-c7d2-4dad-afa0-ccfeada921a8"
				response.ArtistInfo.LastFmUrl = "https://www.last.fm/music/Black+Sabbath"
				response.ArtistInfo.SmallImageUrl = "https://userserve-ak.last.fm/serve/64/27904353.jpg"
				response.ArtistInfo.MediumImageUrl = "https://userserve-ak.last.fm/serve/126/27904353.jpg"
				response.ArtistInfo.LargeImageUrl = "https://userserve-ak.last.fm/serve/_/27904353/Black+Sabbath+sabbath+1970.jpg"
				response.ArtistInfo.SimilarArtist = []Artist{
					{Id: "22", Name: "Accept"},
					{Id: "101", Name: "Bruce Dickinson"},
					{Id: "26", Name: "Aerosmith"},
				}
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})

		})
	})

	Describe("TopSongs", func() {
		BeforeEach(func() {
			response.TopSongs = &TopSongs{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.TopSongs.Song = child
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("SimilarSongs", func() {
		BeforeEach(func() {
			response.SimilarSongs = &SimilarSongs{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.SimilarSongs.Song = child
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("SimilarSongs2", func() {
		BeforeEach(func() {
			response.SimilarSongs2 = &SimilarSongs2{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.SimilarSongs2.Song = child
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("PlayQueue", func() {
		BeforeEach(func() {
			response.PlayQueue = &PlayQueue{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				response.PlayQueue.Username = "user1"
				response.PlayQueue.Current = "111"
				response.PlayQueue.Position = 243
				response.PlayQueue.Changed = &time.Time{}
				response.PlayQueue.ChangedBy = "a_client"
				child := make([]Child, 1)
				child[0] = Child{Id: "1", Title: "title", IsDir: false}
				response.PlayQueue.Entry = child
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Bookmarks", func() {
		BeforeEach(func() {
			response.Bookmarks = &Bookmarks{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				bmk := Bookmark{
					Position: 123,
					Username: "user2",
					Comment:  "a comment",
					Created:  time.Time{},
					Changed:  time.Time{},
				}
				bmk.Entry = Child{Id: "1", Title: "title", IsDir: false}
				response.Bookmarks.Bookmark = []Bookmark{bmk}
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("ScanStatus", func() {
		BeforeEach(func() {
			response.ScanStatus = &ScanStatus{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				t, _ := time.Parse(time.RFC822, time.RFC822)
				response.ScanStatus = &ScanStatus{
					Scanning:    true,
					FolderCount: 123,
					Count:       456,
					LastScan:    &t,
				}
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})
	})

	Describe("Lyrics", func() {
		BeforeEach(func() {
			response.Lyrics = &Lyrics{}
		})

		Context("without data", func() {
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})
		})

		Context("with data", func() {
			BeforeEach(func() {
				response.Lyrics.Artist = "Rick Astley"
				response.Lyrics.Title = "Never Gonna Give You Up"
				response.Lyrics.Value = `Never gonna give you up
				Never gonna let you down
				Never gonna run around and desert you
				Never gonna say goodbye`
			})
			It("should match .XML", func() {
				Expect(xml.Marshal(response)).To(MatchSnapshot())
			})
			It("should match .JSON", func() {
				Expect(json.Marshal(response)).To(MatchSnapshot())
			})

		})
	})
})
