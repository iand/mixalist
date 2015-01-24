package db

import (
    "github.com/iand/mixalist/pkg/playlist"
)

// Get the tags associated with a playlist.
func (db *Database) GetPlaylistTags(pid playlist.PlaylistID) (tags []string, err error) {
    rows, err := db.getQueryable().Query("select tag from mix_playlist_tag where pid = $1", pid)
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

func (db *Database) AddPlaylistTags(pid playlist.PlaylistID, tags ...string) (err error) {
    if db.tx.tx == nil {
        return wrapError(1, NotInTransactionError)
    }
    
    existingTags, err := db.GetPlaylistTags(pid)
    if err != nil {
        db.RollbackTx()
        return err
    }
    
    loop:
    for _, tag := range tags {
        for _, existingTag := range existingTags {
            if tag == existingTag {
                continue loop
            }
        }
        
        _, err = db.tx.Exec("insert into mix_playlist_tag values ($1, $2)", pid, tag)
        if err != nil {
            db.RollbackTx()
            return err
        }
    }
    
    return nil
}
