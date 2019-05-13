package main

import (
	"flag"
	"fmt"
	"os"

	driver "github.com/tjmaynes/learning-golang/driver"
	"github.com/tjmaynes/learning-golang/server"
)

func main() {
	var dbSource = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	var dbType = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
	var serverPort = flag.String("SERVER_PORT", os.Getenv("SERVER_PORT"), "Port to run server on.")

	flag.Parse()

	dbConn, err := driver.ConnectDB(*dbSource, *dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	server.Run(dbConn, *serverPort)
}
