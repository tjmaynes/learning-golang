package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// ConnectDB ..
func ConnectDB(dbType, dbSource string) (*sql.DB, error) {
	db, error := sql.Open(dbType, dbSource)
	if error != nil {
		return nil, error
	}

	if err := db.Ping(); err != nil {
		fmt.Println(dbType, dbSource)
		return nil, err
	}

	return db, nil
}
