package playlist

import (
	"github.com/iand/mixalist/pkg/blobstore"
	"time"
)

type PlaylistID int
type EntryID int
type UserID int

type Playlist struct {
	Pid       PlaylistID
	Title     string
	Owner     *User
	Stars     int
	Tags      []string
	Entries   []*Entry
	ParentPid PlaylistID
	ImageBlobID blobstore.ID
}

type Entry struct {
	Eid      EntryID
	Title    string
	Artist   string
	Album    string
	Duration time.Duration
	SrcName  string	// "youtube" or "soundcloud"
	SrcID    string
	ImageBlobID blobstore.ID
}

type User struct {
	Uid  UserID
	Name string
}
