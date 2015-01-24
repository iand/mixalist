package db

import (
    "database/sql"
    "fmt"
    "runtime"
)

type debugWrappedError struct {
    err error
    file string
    line int
}

func (err debugWrappedError) Error() string {
    //return fmt.Sprintf("[%s:%d] %#v", err.file, err.line, err.err)
    return fmt.Sprintf("[%s:%d] %s", err.file, err.line, err.err.Error())
}

// level 0 = where wrapError is called
// level 1 = where the caller of wrapError is called
// etc.
func wrapError(level int, err error) error {
    if err != nil {
        _, file, line, ok := runtime.Caller(level + 1)
        if ok {
            err = debugWrappedError{
                err: err,
                file: file,
                line: line,
            }
        }
    }
    return err
}

type debugWrappedDB struct {
    db *sql.DB
}

func (d debugWrappedDB) Begin() (debugWrappedTx, error) {
    tx, err := d.db.Begin()
    return debugWrappedTx{tx}, wrapError(1, err)
}

func (d debugWrappedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
    result, err := d.db.Exec(query, args...)
    return result, wrapError(1, err)
}

func (d debugWrappedDB) Query(query string, args ...interface{}) (debugWrappedRows, error) {
    rows, err := d.db.Query(query, args...)
    return debugWrappedRows{rows}, wrapError(1, err)
}

func (d debugWrappedDB) QueryRow(query string, args ...interface{}) debugWrappedRow {
    return debugWrappedRow{d.db.QueryRow(query, args...)}
}

type debugWrappedRow struct {
    row *sql.Row
}

func (d debugWrappedRow) Scan(dest ...interface{}) error {
    return wrapError(1, d.row.Scan(dest...))
}

type debugWrappedRows struct {
    rows *sql.Rows
}

func (d debugWrappedRows) Close() error {
    return wrapError(1, d.rows.Close())
}

func (d debugWrappedRows) Columns() ([]string, error) {
    cols, err := d.rows.Columns()
    return cols, wrapError(1, err)
}

func (d debugWrappedRows) Err() error {
    return wrapError(1, d.rows.Err())
}

func (d debugWrappedRows) Next() bool {
    return d.rows.Next()
}

func (d debugWrappedRows) Scan(dest ...interface{}) error {
    return wrapError(1, d.rows.Scan(dest...))
}

type debugWrappedTx struct {
    tx *sql.Tx
}

func (d debugWrappedTx) Commit() error {
    return wrapError(1, d.tx.Commit())
}

func (d debugWrappedTx) Exec(query string, args ...interface{}) (sql.Result, error) {
    result, err := d.tx.Exec(query, args...)
    return result, wrapError(1, err)
}

func (d debugWrappedTx) Query(query string, args ...interface{}) (debugWrappedRows, error) {
    rows, err := d.tx.Query(query, args...)
    return debugWrappedRows{rows}, wrapError(1, err)
}

func (d debugWrappedTx) QueryRow(query string, args ...interface{}) debugWrappedRow {
    return debugWrappedRow{d.tx.QueryRow(query, args...)}
}

func (d debugWrappedTx) Rollback() error {
    return wrapError(1, d.tx.Rollback())
}
