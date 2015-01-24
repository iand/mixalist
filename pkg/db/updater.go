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
func (db *Database) doUpdates() (empty bool, err error) {
    if db.tx.tx == nil {
        return false, wrapError(1, NotInTransactionError)
    }
    
    current, empty, err := db.getVersion()
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
                err = db.doUpdate(update)
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
func (db *Database) doUpdate(update *DatabaseUpdate) (err error) {
    if db.tx.tx == nil {
        return wrapError(1, NotInTransactionError)
    }
    
    if db.logging {
        log.Printf("Performing database update from version %d to version %d", update.From, update.To)
    }
    
    for _, stmt := range update.SQL {
        _, err = db.tx.Exec(stmt)
        if err != nil {
            db.tx.Rollback()
            db.tx.tx = nil
            return err
        }
    }
    
    _, err = db.tx.Exec("update mix_version set version = $1", update.To)
    if err != nil {
        db.RollbackTx()
        return err
    }
    
    return nil
}

// Get the version currently stored in the database, setting the version to the
// latest if it is not set. The version number and a flag indicating whether the
// database was previously empty are returned.
func (db *Database) getVersion() (v DatabaseVersion, empty bool, err error) {
    if db.tx.tx == nil {
        return 0, false, wrapError(1, NotInTransactionError)
    }
    
    err = db.getQueryable().QueryRow("select version from mix_version").Scan(&v)
    if err != nil {
        if isNonexistentTableError(err) {
            // version table does not exist -> database is empty
            if db.logging {
                log.Printf("Creating table mix_version")
            }
            _, err = db.tx.Exec("create table mix_version (version integer)")
            if err != nil {
                db.RollbackTx()
                return 0, false, err
            }
            _, err = db.tx.Exec("insert into mix_version values ($1)", Latest)
            if err != nil {
                db.RollbackTx()
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
            _, err = db.tx.Exec("insert intro mix_version values ($1)", Latest)
            if err != nil {
                db.RollbackTx()
                return 0, false, err
            }
            return Latest, true, nil
        } else {
            // some other error
            db.tx.Rollback()
            db.tx.tx = nil
            return 0, false, err
        }
    }
    
    // version fetch was successful
    return v, false, nil
}
