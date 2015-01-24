package db

import (
    "github.com/iand/mixalist/pkg/playlist"
)

// Get the number of stars a playlist has.
func (db *Database) GetPlaylistStars(pid playlist.PlaylistID) (stars int, err error) {
    row := db.getQueryable().QueryRow("select count(*) from mix_playlist_star where pid = $1", pid)
    err = row.Scan(&stars)
    if err != nil {
        if isNoRowsError(err) {
            err = InvalidPlaylistError
        }
        return 0, err
    }
    
    return stars, nil
}
