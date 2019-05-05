package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	db "github.com/tjmaynes/learning-golang/db"
)

func main() {
	var dbSource = flag.String("db.source", os.Getenv("DB_SOURCE"), "Database url connection string.")
	var dbType = flag.String("db.type", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
	var migrationDir = flag.String("migration.files", "./db/migrations", "Directory where the migration files are located ?")

	flag.Parse()

	dbConn, err := db.ConnectDB(*dbSource, *dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	log.Println(fmt.Sprintf("Connected to %s db!", *dbType))

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir),
		*dbType, dbConn)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated!")

	os.Exit(0)
}
