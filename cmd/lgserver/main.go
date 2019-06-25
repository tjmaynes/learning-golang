package main

import (
	"flag"
	"os"

	app "github.com/tjmaynes/learning-golang/app"
)

func main() {
	var (
		dbSource   = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
		dbType     = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
		serverPort = flag.String("SERVER_PORT", os.Getenv("SERVER_PORT"), "Port to run server on.")
	)

	flag.Parse()

	a := app.NewApp(*dbSource, *dbType)
	a.Run(*serverPort)
}
