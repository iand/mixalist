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

func (db *Database) ToggleStar(uid playlist.UserID, pid playlist.PlaylistID) (bool, error) {
	if db.tx.tx == nil {
		return false, wrapError(1, NotInTransactionError)
	}

	var count int

	row := db.getQueryable().QueryRow("select count(*) from mix_playlist_star where pid = $1 and uid = $2", pid, uid)
	err := row.Scan(&count)
	if err != nil {
		// something bad happpened
		return false, err
	}

	if count == 0 {
		// No star, so add one
		_, err = db.tx.Exec("insert into mix_playlist_star values ($1, $2)", pid, uid)
		if err != nil {
			db.RollbackTx()
			return false, err
		}
		return true, nil

	} else {
		// We have star, remove it
		_, err = db.tx.Exec("delete from mix_playlist_star where pid = $1 and uid = $2", pid, uid)
		if err != nil {
			db.RollbackTx()
			return false, err
		}
		return false, nil
	}

	return false, nil

}
