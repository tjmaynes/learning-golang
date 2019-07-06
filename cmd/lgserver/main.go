package main

import (
	"flag"
	"os"

	api "github.com/tjmaynes/learning-golang/pkg/api"
)

func main() {
	var (
		dbType     = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
		dbSource   = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database source such as ./db/my.db.")
		serverPort = flag.String("SERVER_PORT", os.Getenv("SERVER_PORT"), "Port to run server from.")
	)

	flag.Parse()

	api.
		NewAPI(*dbType, *dbSource).
		Run(*serverPort)
}
