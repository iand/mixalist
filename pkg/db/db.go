package db

import (
	"database/sql"
	"log"
    
    _ "github.com/lib/pq"
)

type Table struct {
	Name string
	Schema string
}

type Database struct {
	conn debugWrappedDB
	logging bool
}

// Connect to the database.
func Connect(logging bool) (db Database, err error) {
	if logging {
		log.Printf("Connecting to database")
	}
	
	conn, err := sql.Open(driver, source)
	if err != nil {
		return Database{}, err
	}
	
	db = Database{debugWrappedDB{conn}, logging}
	
	err = db.init()
	if err != nil {
		return Database{}, err
	}
	
	return db, nil
}

func (db Database) init() (err error) {
	// Perform entire initialisation in a transaction, to ensure an error partway
	// through does not leave the DB half-initialised.
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	
	_, err = tx.Exec("SET search_path TO mixalist,public")
	if err != nil {
		tx.Rollback()
		return err
	}
	
	empty, err := db.doUpdates(tx)
	if err != nil {
		return err
	}
	
	if empty {
		err = db.createTables(tx)
		if err != nil {
			return err
		}
	}
	
	err = tx.Commit()
	if err != nil {
		return err
	}
	
	if db.logging {
		log.Printf("Database initialisation complete")
	}
	
	return nil
}

func (db Database) createTables(tx debugWrappedTx) (err error) {
	for _, table := range LatestSchema {
		if db.logging {
			log.Printf("Creating table %s", table.Name)
		}
		
		_, err = tx.Exec(table.Schema)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	return nil
}
