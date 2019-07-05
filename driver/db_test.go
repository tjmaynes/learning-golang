package db

import (
	"flag"
	"os"
	"testing"
)

var (
	dbSource = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	dbType   = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
)

func Test_Driver_DB_ConnectDB_ShouldReturnDatabaseConnection(t *testing.T) {
	// result, err := ConnectDB(*dbSource, *dbType)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println(*result)
}
