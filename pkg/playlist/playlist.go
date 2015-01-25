package playlist

import (
	"fmt"
	"github.com/iand/mixalist/pkg/blobstore"
	"strings"
	"time"
)

type PlaylistID int
type EntryID int
type UserID int

type Playlist struct {
	Pid         PlaylistID
	Title       string
	Owner       *User
	Stars       int
	Tags        []string
	Entries     []*Entry
	ParentPid   PlaylistID
	ImageBlobID blobstore.ID
}

type Entry struct {
	Eid         EntryID
	Title       string
	Artist      string
	Album       string
	Duration    time.Duration
	SrcName     string // "youtube" or "soundcloud"
	SrcID       string
	ImageBlobID blobstore.ID
}

type User struct {
	Uid  UserID
	Name string
}

type artistCount struct {
	artist string
	count  int
}

func (p *Playlist) Featuring() string {
	switch len(p.Entries) {
	case 0:
		return "No songs :("
	case 1:
		return fmt.Sprintf("Featuring %s", p.Entries[0].Artist)
	}

	artists := map[string]int{}
	for _, e := range p.Entries {
		artists[strings.TrimSpace(e.Artist)]++
	}

	artistCounts := []artistCount{}
	for artist, count := range artists {
		if len(artistCounts) < 3 {
			artistCounts = append(artistCounts, artistCount{
				artist: artist,
				count:  count,
			})
			continue
		}
		for i := 0; i < len(artistCounts); i++ {
			if artistCounts[i].count < count {
				artistCounts[i].artist = artist
				artistCounts[i].count = count
				continue
			}
		}
	}

	switch len(artistCounts) {
	case 1:
		return fmt.Sprintf("Featuring %s", artistCounts[0].artist)
	case 2:
		return fmt.Sprintf("Featuring %s and %s", artistCounts[0].artist, artistCounts[1].artist)
	default:
		return fmt.Sprintf("Featuring %s, %s and %s", artistCounts[0].artist, artistCounts[1].artist, artistCounts[2].artist)
	}

	return "No songs :("

}
