package db

import (
    "github.com/iand/mixalist/pkg/playlist"
)

// Get only the information stored in the actual mix_user record.
func (db Database) GetUserInfo(uid playlist.UserID) (name string, err error) {
    row := db.conn.QueryRow("SELECT name FROM mix_user WHERE uid = $1", uid)
    err = row.Scan(&name)
    if err != nil {
        if isNoRowsError(err) {
            err = InvalidUserError
        }
        return "", err
    }
    
    return name, nil
}

// Get all information about a user.
func (db Database) GetUser(uid playlist.UserID) (user *playlist.User, err error) {
    name, err := db.GetUserInfo(uid)
    if err != nil {
        return nil, err
    }
    
    user = &playlist.User{
        Uid: uid,
        Name: name,
    }
    
    return user, nil
}
