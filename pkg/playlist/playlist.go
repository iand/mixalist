package playlist

import (
    "time"
)

type PlaylistID int
type EntryID int
type UserID int

type Playlist struct {
    Pid     PlaylistID
    Title   string
    Owner   *User
    Stars   int
    Tags    []string
    Entries []*Entry
    ParentPid PlaylistID
}

type Entry struct {
    Eid      EntryID
    Ytid     string
    Title    string
    Artist   string
    Album    string
    Duration time.Duration
}

type User struct {
    Uid  UserID
    Name string
}
