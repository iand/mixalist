package db

import (
    "github.com/iand/mixalist/pkg/playlist"
    "time"
)

// Gets the entries in a playlist.
func (db Database) GetPlaylistEntries(pid playlist.PlaylistID) (entries []*playlist.Entry, err error) {
    rows, err := db.conn.Query("SELECT eid, index, yt_id, title, artist, album, duration FROM mix_playlist_entry WHERE pid = $1", pid)
    if err != nil {
        return nil, err
    }
    
    for rows.Next() {
        var eid playlist.EntryID
        var index, duration int
        var ytid, title, artist, album string
        
        err = rows.Scan(&eid, &index, &ytid, &title, &artist, &album, &duration)
        if err != nil {
            return nil, err
        }
        
        entry := &playlist.Entry{
            Eid: eid,
            Ytid: ytid,
            Title: title,
            Artist: artist,
            Album: album,
            Duration: time.Duration(duration) * time.Second,
        }
        
        if index >= len(entries) {
            entries2 := make([]*playlist.Entry, index, index*2)
            copy(entries2, entries)
            entries = entries2
        }
        
        entries[index] = entry
    }
    
    err = rows.Err()
    if err != nil {
        return nil, err
    }
    
    return entries, nil
}
