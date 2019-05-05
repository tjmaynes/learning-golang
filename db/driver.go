package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB ..
type DB struct {
	*sql.DB
}

// ConnectDB ..
func ConnectDB(dbSource, dbType string) (*DB, error) {
	d, error := sql.Open(dbType, dbSource)
	if error != nil {
		return nil, error
	}
	defer d.Close()

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return &DB{d}, nil
}
