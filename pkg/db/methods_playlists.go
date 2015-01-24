package db

import (
    "github.com/iand/mixalist/pkg/playlist"
)

// Get only the information stored in the actual mix_playlist record.
func (db Database) GetPlaylistInfo(pid playlist.PlaylistID) (title string, ownerUid playlist.UserID, err error) {
    row := db.conn.QueryRow("SELECT title, owner_uid FROM mix_playlist WHERE pid = $1", pid)
    err = row.Scan(&title, &ownerUid)
    if err != nil {
        if isNoRowsError(err) {
            err = InvalidPlaylistError
        }
        return "", 0, err
    }
    
    return title, ownerUid, nil
}

// Get all information about a playlist.
func (db Database) GetPlaylist(pid playlist.PlaylistID) (p *playlist.Playlist, err error) {
    title, ownerUid, err := db.GetPlaylistInfo(pid)
    if err != nil {
        return nil, err
    }
    
    owner, err := db.GetUser(ownerUid)
    if err != nil {
        return nil, err
    }
    
    stars, err := db.GetPlaylistStars(pid)
    if err != nil {
        return nil, err
    }
    
    tags, err := db.GetPlaylistTags(pid)
    if err != nil {
        return nil, err
    }
    
    entries, err := db.GetPlaylistEntries(pid)
    if err != nil {
        return nil, err
    }
    
    return &playlist.Playlist{
        Pid: pid,
        Title: title,
        Owner: owner,
        Stars: stars,
        Tags: tags,
        Entries: entries,
    }, nil
}
