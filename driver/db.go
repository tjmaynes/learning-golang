package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConnectDB ..
func ConnectDB(dbSource, dbType string) (*sql.DB, error) {
	db, error := sql.Open(dbType, dbSource)
	if error != nil {
		return nil, error
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
