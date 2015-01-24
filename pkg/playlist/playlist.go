package playlist

import (
	"time"
)

type Playlist struct {
	Pid     int
	Title   string
	Owner   User
	Stars   int
	Tags    []string
	Entries []Entry
}

type Entry struct {
	Eid      int
	Ytid     string
	Title    string
	Artist   string
	Album    string
	Duration time.Duration
}

type User struct {
	Uid  int
	Name string
}
