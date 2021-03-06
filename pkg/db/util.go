package db

import (
    "database/sql"
    "errors"
    "github.com/lib/pq"
    "strings"
)

var (
    InvalidPlaylistError  = errors.New("Invalid playlist ID")
    InvalidUserError      = errors.New("Invalid user ID")
    NotInTransactionError = errors.New("Attempt to modify database where not in a transaction")
)

// http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
const invalidTableErrorCode pq.ErrorCode = "42P01"

func isNonexistentTableError(err error) bool {
    dwe, ok := err.(debugWrappedError)
    if ok {
        err = dwe.err
    }
    perr, ok := err.(*pq.Error)
    return ok && perr.Code == invalidTableErrorCode
}

func isNoRowsError(err error) bool {
    dwe, ok := err.(debugWrappedError)
    if ok {
        err = dwe.err
    }
    return err == sql.ErrNoRows
}

// escape string to go into an SQL LIKE pattern
func patternEscape(s string) string {
    s = strings.Replace(s, "\\", "\\\\", -1)
    s = strings.Replace(s, "%", "\\%", -1)
    s = strings.Replace(s, "_", "\\_", -1)
    return s
}
