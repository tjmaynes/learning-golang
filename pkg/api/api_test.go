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
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/icrowley/fake"
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

func Test_CartEndpoint_WhenUnsupportedMethodIsGiven_Returns405(t *testing.T) {
	flag.Parse()

	var tests = []struct {
		httpMethod string
		endpoint   string
	}{
		{"PUT", "/cart"},
		{"PUT", "/cart/123"},
	}

	for _, tt := range tests {
		a := NewAPI(*dbSource, *dbType)

		request, err := http.NewRequest(tt.httpMethod, tt.endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		a.Handler.ServeHTTP(recorder, request)

		if http.StatusMethodNotAllowed != recorder.Code {
			t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, recorder.Code)
		}
	}
}

func Test_CartEndpoint_GetAllItems_WhenItemsExist_ShouldReturnAllItems(t *testing.T) {
	flag.Parse()

	a := NewAPI(*dbSource, *dbType)

	ctx := context.Background()
	cartRepository := cart.NewRepository(getDbConn())
	items := setupDatabase(ctx, cartRepository)

	limit := 5
	requestURL := fmt.Sprintf("/cart?limit=%d", limit)
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Handler.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, recorder.Code)
	}

	expected := createResponseBody(items[:limit])

	if body := recorder.Body.String(); body != expected {
		t.Errorf("Expected an array of cart items. Got %s", body)
	}

	teardownDatabase(ctx, cartRepository, items)
}

func Test_CartEndpoint_GetItemByID_WhenItemExists_ShouldReturnItem(t *testing.T) {
	flag.Parse()

	a := NewAPI(*dbSource, *dbType)

	ctx := context.Background()
	cartRepository := cart.NewRepository(getDbConn())
	items := setupDatabase(ctx, cartRepository)

	item1 := items[0]
	requestURL := fmt.Sprintf("/cart/%d", item1.ID)

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Handler.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, recorder.Code)
	}

	expected := createResponseBody(item1)

	if body := recorder.Body.String(); body != expected {
		t.Errorf("Expected an array of cart items. Got %s", body)
	}

	teardownDatabase(ctx, cartRepository, items)
}

func Test_CartEndpoint_AddItem_WhenGivenValidItem_ShouldReturnItem(t *testing.T) {
	flag.Parse()

	a := NewAPI(*dbSource, *dbType)

	ctx := context.Background()
	cartRepository := cart.NewRepository(getDbConn())

	itemName := fake.ProductName()
	itemPrice := cart.Decimal(99)
	itemManufacturer := fake.Brand()
	newItem := cart.Item{Name: itemName, Price: itemPrice, Manufacturer: itemManufacturer}

	form := url.Values{}
	form.Add("name", newItem.Name)
	form.Add("price", fmt.Sprintf("%d", newItem.Price))
	form.Add("manufacturer", newItem.Manufacturer)

	request, err := http.NewRequest("POST", "/cart", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	a.Handler.ServeHTTP(recorder, request)

	if http.StatusCreated != recorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, recorder.Code)
	}

	var result struct {
		Data cart.Item `json:"data"`
	}
	err = json.Unmarshal([]byte(recorder.Body.String()), &result)
	if err != nil {
		t.Fatal(err)
	}

	newItem.ID = result.Data.ID

	if result.Data != newItem {
		t.Errorf("Expected a cart item %+v. Got %+v", newItem, result.Data)
	}

	items := []cart.Item{newItem}
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
