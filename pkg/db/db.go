package db

import (
	"database/sql"
)

type Database struct {
	conn sql.DB
}

// Connect to the database.
func Connect() (db Database, err error) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		return Database{}, err
	}
	
	db = Database{Conn}
	err = db.doUpdates()
	if err != nil {
		return Database{}, err
	}
	
	return db, nil
}
