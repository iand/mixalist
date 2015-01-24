package db

import (
    "fmt"
    "log"
)

type DatabaseVersion uint

type DatabaseUpdate struct {
    From, To DatabaseVersion
    SQL []string
}

// Check the database version and perform updates if necessary. A flag indicating
// whether the database was previously empty is returned.
func (db Database) doUpdates(tx debugWrappedTx) (empty bool, err error) {
    current, empty, err := db.getVersion(tx)
    if err != nil {
        return false, err
    }
    
    if db.logging {
        log.Printf("Current database version is %d, latest is %d", current, Latest)
    }
    
    for current < Latest {
        updateDone := false
        
        for _, update := range Updates {
            if update.From == current {
                err = db.doUpdate(tx, update)
                if err != nil {
                    return false, err
                }
                
                current = update.To
                updateDone = true
                break
            }
        }
        
        if !updateDone {
            return false, fmt.Errorf("Failed to update database from current version %d to latest version %d (no suitable update found)", current, Latest)
        }
    }
    
    if db.logging {
        log.Printf("Database is up to date")
    }
    
    return empty, nil
}

// Execute a DatabaseUpdate and update the version number stored in the database.
func (db Database) doUpdate(tx debugWrappedTx, update *DatabaseUpdate) (err error) {
    if db.logging {
        log.Printf("Performing database update from version %d to version %d", update.From, update.To)
    }
    
    for _, stmt := range update.SQL {
        _, err = tx.Exec(stmt)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    
    _, err = tx.Exec("UPDATE mix_version SET version = $1", update.To)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    return nil
}

// Get the version currently stored in the database, setting the version to the
// latest if it is not set. The version number and a flag indicating whether the
// database was previously empty are returned.
func (db Database) getVersion(tx debugWrappedTx) (v DatabaseVersion, empty bool, err error) {
    err = db.conn.QueryRow("SELECT version FROM mix_version").Scan(&v)
    if err != nil {
        if isNonexistentTableError(err) {
            // version table does not exist -> database is empty
            if db.logging {
                log.Printf("Creating table mix_version")
            }
            _, err = tx.Exec("CREATE TABLE mix_version (version integer)")
            if err != nil {
                tx.Rollback()
                return 0, false, err
            }
            _, err = tx.Exec("INSERT INTO mix_version VALUES ($1)", Latest)
            if err != nil {
                tx.Rollback()
                return 0, false, err
            }
            return Latest, true, nil
        } else if isNoRowsError(err) {
            // version table exists and is empty -> assume database empty
            // This shouldn't happen unless the database is corrupted, in which
            // case corrupting it further is not that much of an issue.
            if db.logging {
                log.Printf("Version table exists but is empty - was the database in a corrupted state?")
            }
            _, err = tx.Exec("INSERT INTO mix_version VALUES ($1)", Latest)
            if err != nil {
                tx.Rollback()
                return 0, false, err
            }
            return Latest, true, nil
        } else {
            // some other error
            tx.Rollback()
            return 0, false, err
        }
    }
    
    // version fetch was successful
    return v, false, nil
}
