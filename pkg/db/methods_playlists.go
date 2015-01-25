package db

import (
	"fmt"
	"github.com/iand/mixalist/pkg/playlist"
	"strings"
)

// $1 - pagination start index
// $2 - pagination page size
// $3... - tags to match
const scoringQueryHeader = `
    with stars as (
        select pid, count(*)
        from mix_playlist_star
        group by pid
    )
    select mix_playlist.pid as p, (coalesce(stars.count, 0) + 1) /
        ((extract(epoch from
            (timestamp 'now' - mix_playlist.created)) / 3600 + 2) ^ 1.8) as score
    from mix_playlist
    left join stars on mix_playlist.pid = stars.pid`
const scoringQueryFooter = `
    order by score desc
    limit $2
    offset $1`

// Sort the playlists by score and return those that have any of the tags specified
// in requiredTags. If requiredTags is empty, all playlists are sorted. The results
// are paginated using the pageSize and pageNum arguments (pageNum is 0-indexed).
func (db *Database) GetSortedPlaylistIDs(pageSize, pageNum int, requiredTags []string) (pids []playlist.PlaylistID, err error) {
	start := pageNum * pageSize

	query := scoringQueryHeader
	params := []interface{}{start, pageSize}
	if len(requiredTags) > 0 {
		query += " where exists(select tag from mix_playlist_tag where pid = p and ("
		for i, tag := range requiredTags {
			if i > 0 {
				query += " or "
			}
			query += fmt.Sprintf("tag = $%d", i+3)
			params = append(params, tag)
		}
		query += "))"
	}
	query += scoringQueryFooter

	rows, err := db.getQueryable().Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pid playlist.PlaylistID
		var score float32
		err = rows.Scan(&pid, &score)
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return pids, nil
}

// Get only the information stored in the actual mix_playlist record.
func (db *Database) GetPlaylistRecord(pid playlist.PlaylistID) (title string, ownerUid playlist.UserID, parentPid playlist.PlaylistID, err error) {
	row := db.getQueryable().QueryRow("select title, owner_uid, coalesce(parent_pid, 0) from mix_playlist where pid = $1", pid)
	err = row.Scan(&title, &ownerUid, &parentPid)
	if err != nil {
		if isNoRowsError(err) {
			err = InvalidPlaylistError
		}
		return "", 0, 0, err
	}

	return title, ownerUid, parentPid, nil
}

// Get all information about a playlist.
func (db *Database) GetPlaylist(pid playlist.PlaylistID) (p *playlist.Playlist, err error) {
	title, ownerUid, parentPid, err := db.GetPlaylistRecord(pid)
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
		Pid:       pid,
		Title:     title,
		Owner:     owner,
		Stars:     stars,
		Tags:      tags,
		Entries:   entries,
		ParentPid: parentPid,
	}, nil
}

func (db *Database) CreatePlaylistRecord(title string, ownerUid playlist.UserID, parentPid playlist.PlaylistID, searchText string) (newPid playlist.PlaylistID, err error) {
	if db.tx.tx == nil {
		return 0, wrapError(1, NotInTransactionError)
	}

	row := db.tx.QueryRow("insert into mix_playlist (title, owner_uid, created, search_text, parent_pid) values ($1, $2, timestamp 'now', $3, $4) returning pid", title, ownerUid, searchText, parentPid)
	err = row.Scan(&newPid)
	if err != nil {
		db.RollbackTx()
		return 0, err
	}
	return newPid, nil
}

// Create a new playlist with the given data. p.Pid and p.Entries[i].Eid are set
// to their newly assigned values.
func (db *Database) CreatePlaylist(p *playlist.Playlist) (newPid playlist.PlaylistID, err error) {
	if db.tx.tx == nil {
		return 0, wrapError(1, NotInTransactionError)
	}

	searchText := p.Title
	for _, tag := range p.Tags {
		searchText += " " + tag
	}
	searchText = strings.ToLower(searchText)

	pid, err := db.CreatePlaylistRecord(p.Title, p.Owner.Uid, p.ParentPid, searchText)
	if err != nil {
		return 0, err
	}
	p.Pid = pid

	err = db.AddPlaylistTags(pid, p.Tags...)
	if err != nil {
		return 0, err
	}

	for i, entry := range p.Entries {
		eid, err := db.CreatePlaylistEntry(i, pid, entry)
		if err != nil {
			return pid, err
		}

		entry.Eid = eid
	}

	return pid, nil
}

func (db *Database) SearchPlaylists(pageSize, pageNum int, keywords ...string) (pids []playlist.PlaylistID, err error) {
	start := pageNum * pageSize
	params := []interface{}{start, pageSize}
	query := "select pid from mix_playlist"

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
		var pid playlist.PlaylistID
		err = rows.Scan(&pid)
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return pids, nil
}
