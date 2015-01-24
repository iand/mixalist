package db

import (
    "github.com/iand/mixalist/pkg/playlist"
)

// Get the tags associated with a playlist.
func (db Database) GetPlaylistTags(pid playlist.PlaylistID) (tags []string, err error) {
    rows, err := db.conn.Query("SELECT tag FROM mix_playlist_tag WHERE pid = $1", pid)
    if err != nil {
        return nil, err
    }
    
    for rows.Next() {
        var tag string
        err = rows.Scan(&tag)
        if err != nil {
            return nil, err
        }
        tags = append(tags, tag)
    }
    
    err = rows.Err()
    if err != nil {
        return nil, err
    }
    
    return tags, nil
}

func (db Database) AddPlaylistTags(tx debugWrappedTx, pid playlist.PlaylistID, tags ...string) (err error) {
    existingTags, err := db.GetPlaylistTags(pid)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    loop:
    for _, tag := range tags {
        for _, existingTag := range existingTags {
            if tag == existingTag {
                continue loop
            }
        }
        
        _, err = db.conn.Exec("INSERT INTO mix_playlist_tag VALUES ($1, $2)", pid, tag)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    
    return nil
}
