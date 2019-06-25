package app

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	dbSource = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	dbType   = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
)

func setupDatabase() {

}

func teardownDatabase() {

}

func Test_PingEndpoint_ReturnsPong(t *testing.T) {
	flag.Parse()

	a := NewApp(*dbSource, *dbType)

	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	if body := response.Body.String(); body != `{"message":"PONG!"}` {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

// func Test_PostEndpoint_Post
