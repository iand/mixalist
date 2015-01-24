package db

type DatabaseVersion uint

type DatabaseUpdate struct {
    From, To DatabaseVersion
    SQL []string
}

// Check the database version and perform updates if necessary.
func (db Database) doUpdates() (err error) {
    current, err := db.GetVersion()
    if err != nil {
        return err
    }
    
    for current < Latest {
        updateDone := false
        
        for _, update := range Updates {
            if update.From == current {
                err = db.doUpdate(update)
                if err != nil {
                    return err
                }
                
                current = update.To
                updateDone = true
                break
            }
        }
        
        if !updateDone {
            return fmt.Errorf("Failed to update database from current version %d to latest version %d (no suitable update found)", current, Latest)
        }
    }
    
    return nil
}

// Execute a DatabaseUpdate and update the version number stored in the database.
func (db Database) doUpdate(update *DatabaseUpdate) (err error) {
    tx, err := db.conn.Begin()
    if err != nil {
        return err
    }
    
    for _, stmt := range update.SQL {
        _, err = tx.Exec(stmt)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    
    _, err = tx.Exec("UPDATE version SET version = ?", update.To)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    return tx.Commit()
}

// Get the version currently stored in the database, setting the version to the
// latest if it is not set.
func (db Database) GetVersion() (v DatabaseVersion, err error) {
    err = db.conn.QueryRow("SELECT version FROM version").Scan(&v)
    if err != nil {
        isNTE := isNonexistentTableError(err)
        if isNTE {
            // version table does not exist -> database is empty
            err = db.conn.Exec("CREATE TABLE version (version integer)")
            if err != nil {
                return 0, err
            }
            err = db.conn.Exec("INSERT INTO version VALUES (?)", Latest)
            if err != nil {
                return 0, err
            }
            return Latest, nil
        } else {
            // some other error
            return 0, err
        }
    }
    
    // version fetch was successful
    return v, nil
}

func isNonexistentTableError(err error) bool {
    mysqlError, isMysqlError := err.(*mysql.MySQLError)
    return isMysqlError && mysqlError.Number == 1146
}
