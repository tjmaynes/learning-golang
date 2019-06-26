package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	driver "github.com/tjmaynes/learning-golang/driver"
	cart "github.com/tjmaynes/learning-golang/pkg/cart"
)

var (
	dbSource   = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	dbType     = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")
	jsonSource = flag.String("JSON_SOURCE", os.Getenv("JSON_SOURCE"), "JSON Source, such as ./cmd/data.json.")
)

func Test_PingEndpoint_ReturnsPong(t *testing.T) {
	flag.Parse()

	a := NewAPI(*dbSource, *dbType)

	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	a.Handler.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, recorder.Code)
	}

	if body := recorder.Body.String(); body != `{"message":"PONG!"}` {
		t.Errorf("Expected a PONG! message. Got %s", body)
	}
}

func Test_CartEndpoint_GetAllItems(t *testing.T) {
	flag.Parse()

	a := NewAPI(*dbSource, *dbType)

	ctx := context.Background()
	cartRepository := cart.NewRepository(getDbConn())
	items := setupDatabase(ctx, cartRepository)

	request, err := http.NewRequest("GET", "/cart", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	a.Handler.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, recorder.Code)
	}

	expected := createResponseBody(items[:10])

	if body := recorder.Body.String(); body != expected {
		t.Errorf("Expected an array of cart items. Got %s", body)
	}

	teardownDatabase(ctx, cartRepository, items)
}

func createResponseBody(items interface{}) string {
	out, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`{"data":%s}`, out)
}

func getDbConn() *sql.DB {
	dbConn, err := driver.ConnectDB(*dbSource, *dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return dbConn
}

func setupDatabase(ctx context.Context, cartRepository cart.Repository) []cart.Item {
	flag.Parse()

	jsonFile, err := os.Open(*jsonSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var rawItems []cart.Item
	err = json.Unmarshal([]byte(jsonBytes), &rawItems)
	if err != nil {
		panic(err)
	}

	var items []cart.Item
	for _, rawItem := range rawItems {
		item, err := cartRepository.AddItem(ctx, &rawItem)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}

	return items
}

func teardownDatabase(ctx context.Context, cartRepository cart.Repository, Items []cart.Item) {
	for _, Item := range Items {
		_, err := cartRepository.RemoveItem(ctx, Item.ID)
		if err != nil {
			panic(err)
		}
	}
}
