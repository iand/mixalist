package db

import (
    "github.com/iand/mixalist/pkg/playlist"
    "time"
)

// Gets the entries in a playlist.
func (db Database) GetPlaylistEntries(pid playlist.PlaylistID) (entries []*playlist.Entry, err error) {
    rows, err := db.conn.Query("SELECT eid, yt_id, title, artist, album, duration FROM mix_playlist_entry WHERE pid = $1 ORDER BY index", pid)
    if err != nil {
        return nil, err
    }
    
    for rows.Next() {
        var eid playlist.EntryID
        var ytid, title, artist, album string
        var duration int
        
        err = rows.Scan(&eid, &ytid, &title, &artist, &album, &duration)
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
        
        entries = append(entries, entry)
    }
    
    err = rows.Err()
    if err != nil {
        return nil, err
    }
    
    return entries, nil
}
