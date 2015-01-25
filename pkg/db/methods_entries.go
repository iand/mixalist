package db

import (
	"fmt"
	"github.com/iand/mixalist/pkg/playlist"
	"time"
)

// Gets the entries in a playlist.
func (db *Database) GetPlaylistEntries(pid playlist.PlaylistID) (entries []*playlist.Entry, err error) {
	rows, err := db.getQueryable().Query("select eid, title, artist, album, duration, src_name, src_id from mix_playlist_entry where pid = $1 order by index", pid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var eid playlist.EntryID
		var title, artist, album, srcName, srcID string
		var duration int

		err = rows.Scan(&eid, &title, &artist, &album, &duration, &srcName, &srcID)
		if err != nil {
			return nil, err
		}

		entry := &playlist.Entry{
			Eid:      eid,
			Title:    title,
			Artist:   artist,
			Album:    album,
			SrcName:  srcName,
			SrcID:    srcID,
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
	row := db.tx.QueryRow("insert into mix_playlist_entry (pid, index, title, artist, album, duration, search_text, src_name, src_id) values ($1, $2, $3, $4, $5, $6, lower($7 || ' ' || $8 || ' ' || $9), $10, $11) returning eid", pid, index, entry.Title, entry.Artist, entry.Album, duration, entry.Title, entry.Artist, entry.Album, entry.SrcName, entry.SrcID)
	err = row.Scan(&newEid)
	if err != nil {
		db.RollbackTx()
		return 0, err
	}
	return newEid, nil
}

func (db *Database) SearchEntries(pageSize, pageNum int, keywords ...string) (entries []*playlist.Entry, err error) {
	start := pageNum * pageSize
	params := []interface{}{start, pageSize}
	query := "select eid, title, artist, album, duration, src_name, src_id from mix_playlist_entry"

	for i, keyword := range keywords {
		params = append(params, "%"+patternEscape(keyword)+"%")
		if i > 0 {
			query += " and "
		} else {
			query += " where "
		}
		query += fmt.Sprintf("search_text like $%d", len(params))
	}

	query += " limit $2 offset $1"
	rows, err := db.getQueryable().Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var duration int
		entry := new(playlist.Entry)
		err = rows.Scan(&entry.Eid, &entry.Title, &entry.Artist, &entry.Album, &duration, &entry.SrcName, &entry.SrcID)
		if err != nil {
			return nil, err
		}
		entry.Duration = time.Duration(duration) * time.Second
		entries = append(entries, entry)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return entries, nil
}
