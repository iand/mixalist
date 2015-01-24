package db

import (
    "fmt"
    "github.com/iand/mixalist/pkg/playlist"
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
func (db Database) GetSortedPlaylistIDs(pageSize, pageNum int, requiredTags []string) (pids []playlist.PlaylistID, err error) {
    start := pageNum * pageSize
    
    query := scoringQueryHeader
    params := []interface{}{start, pageSize}
    if len(requiredTags) > 0 {
        query += " where exists(select tag from mix_playlist_tag where pid = p and ("
        for i, tag := range requiredTags {
            if i > 0 {
                query += " or "
            }
            query += fmt.Sprintf("tag = $%d", i + 3)
            params = append(params, tag)
        }
        query += "))"
    }
    query += scoringQueryFooter
    
    rows, err := db.conn.Query(query, params...)
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
func (db Database) GetPlaylistInfo(pid playlist.PlaylistID) (title string, ownerUid playlist.UserID, err error) {
    row := db.conn.QueryRow("SELECT title, owner_uid FROM mix_playlist WHERE pid = $1", pid)
    err = row.Scan(&title, &ownerUid)
    if err != nil {
        if isNoRowsError(err) {
            err = InvalidPlaylistError
        }
        return "", 0, err
    }
    
    return title, ownerUid, nil
}

// Get all information about a playlist.
func (db Database) GetPlaylist(pid playlist.PlaylistID) (p *playlist.Playlist, err error) {
    title, ownerUid, err := db.GetPlaylistInfo(pid)
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
        Pid: pid,
        Title: title,
        Owner: owner,
        Stars: stars,
        Tags: tags,
        Entries: entries,
    }, nil
}
