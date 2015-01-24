package db

import (
	"database/sql"
	"log"
    
    _ "github.com/lib/pq"
)

type queryable interface {
	Query(string, ...interface{}) (debugWrappedRows, error)
	QueryRow(string, ...interface{}) debugWrappedRow
}

type Table struct {
	Name string
	Schema string
}

type Database struct {
	conn debugWrappedDB
	tx debugWrappedTx
	logging bool
}

// Connect to the database.
func Connect(logging bool) (db *Database, err error) {
	if logging {
		log.Printf("Connecting to database")
	}
	
	conn, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	
	db = &Database{
		conn: debugWrappedDB{conn},
		tx: debugWrappedTx{nil},
		logging: logging,
	}
	
	err = db.BeginTx()
	if err != nil {
		return nil, err
	}
	
	err = db.init()
	if err != nil {
		return nil, err
	}
	
	err = db.CommitTx()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (db *Database) BeginTx() (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	db.tx = tx
	return nil
}

func (db *Database) RollbackTx() (err error) {
	err = db.tx.Rollback()
	db.tx.tx = nil
	return err
}

func (db *Database) CommitTx() (err error) {
	err = db.tx.Commit()
	db.tx.tx = nil
	return err
}

func (db *Database) getQueryable() (q queryable) {
	if db.tx.tx != nil {
		return db.tx
	}
	return db.conn
}

func (db *Database) init() (err error) {
	if db.tx.tx == nil {
		return wrapError(1, NotInTransactionError)
	}
	
	_, err = db.tx.Exec("SET search_path TO mixalist,public")
	if err != nil {
		db.RollbackTx()
		return err
	}
	
	empty, err := db.doUpdates()
	if err != nil {
		return err
	}
	
	if empty {
		err = db.createTables()
		if err != nil {
			return err
		}
	}
	
	if db.logging {
		log.Printf("Database initialisation complete")
	}
	
	return nil
}

func (db *Database) createTables() (err error) {
	if db.tx.tx == nil {
		return wrapError(1, NotInTransactionError)
	}
	
	for _, table := range LatestSchema {
		if db.logging {
			log.Printf("Creating table %s", table.Name)
		}
		
		_, err = db.tx.Exec(table.Schema)
		if err != nil {
			db.RollbackTx()
			return err
		}
	}
	
	return nil
}
