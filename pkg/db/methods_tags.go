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
