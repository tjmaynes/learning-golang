package main

import (
	"flag"

	api "github.com/tjmaynes/learning-golang/pkg/api"
)

func main() {
	var (
		dbSource   = flag.String("db-source", "mysql-user:password@/learning-golang-db", "Database url connection string.")
		dbType     = flag.String("db-type", "mysql", "Database Type, such as postgres, mysql, etc.")
		serverPort = flag.String("server-port", "3000", "Port to run server from.")
	)

	flag.Parse()

	api.
		NewAPI(*dbSource, *dbType).
		Run(*serverPort)
}
