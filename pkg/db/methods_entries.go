package db

import (
    "github.com/iand/mixalist/pkg/playlist"
    "time"
)

// Gets the entries in a playlist.
func (db *Database) GetPlaylistEntries(pid playlist.PlaylistID) (entries []*playlist.Entry, err error) {
    rows, err := db.getQueryable().Query("SELECT eid, yt_id, title, artist, album, duration FROM mix_playlist_entry WHERE pid = $1 ORDER BY index", pid)
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

func (db *Database) CreatePlaylistEntry(index int, pid playlist.PlaylistID, entry *playlist.Entry) (newEid playlist.EntryID, err error) {
    if db.tx.tx == nil {
        return 0, wrapError(1, NotInTransactionError)
    }
    
    duration := int(entry.Duration / time.Second)
    row := db.tx.QueryRow("insert into mix_playlist_entry (pid, index, yt_id, title, artist, album, duration, search_text) values ($1, $2, $3, $4, $5, $6, $7, $4 || ' ' || $5 || ' ' || $6) returning eid", pid, index, entry.Ytid, entry.Title, entry.Artist, entry.Album, duration)
    err = row.Scan(&newEid)
    if err != nil {
        db.RollbackTx()
        return 0, err
    }
    return newEid, nil
}
