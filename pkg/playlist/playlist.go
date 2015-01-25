package playlist

import (
	"fmt"
	"github.com/iand/mixalist/pkg/blobstore"
	"math/rand"
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
		if strings.TrimSpace(p.Entries[0].Artist) != "" {
			return fmt.Sprintf("Featuring %s", p.Entries[0].Artist)
		}
		if strings.TrimSpace(p.Entries[0].Title) != "" {
			return fmt.Sprintf("Featuring %s", p.Entries[0].Title)
		}
		return "One song"
	}

	artists := []string{}
	perm := rand.Perm(len(p.Entries))

outer:
	for _, i := range perm {
		if len(artists) > 2 {
			break
		}

		artist := strings.TrimSpace(p.Entries[i].Artist)
		if artist == "" {
			continue
		}

		for _, existing := range artists {
			if strings.EqualFold(strings.TrimSpace(existing), artist) {
				continue outer
			}
		}
		artists = append(artists, artist)
	}

	switch len(artists) {
	case 1:
		return fmt.Sprintf("Featuring %s", artists[0])
	case 2:
		return fmt.Sprintf("Featuring %s and %s", artists[0], artists[1])
	case 3:
		return fmt.Sprintf("Featuring %s, %s and %s", artists[0], artists[1], artists[2])
	}

	return fmt.Sprintf("%d songs", len(p.Entries))

}
