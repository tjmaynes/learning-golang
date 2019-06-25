package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
)

var (
	dbSource   = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	dbType     = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
	serverPort = flag.String("SERVER_PORT", os.Getenv("SERVER_PORT"), "Port to run server on.")
)

func setupDatabase() {

}

func teardownDatabase() {

}

func TestMain(m *testing.M) {
    a = App{}
    a.Initialize("DB_USERNAME", "DB_PASSWORD", "rest_api_example")
    ensureTableExists()
    code := m.Run()
    clearTable()
    os.Exit(code)
}

func Test_HealthCheckEndpoint_ReturnsPong(t *testing.T) {
	flag.Parse()

	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	os.Exit(1)
}
