package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	driver "github.com/tjmaynes/learning-golang/driver"
	cart "github.com/tjmaynes/learning-golang/pkg/cart"
)

// SeedData ..
func SeedData(jsonSource string, dbConn *sql.DB) []int64 {
	cartRepository := cart.NewRepository(dbConn)
	ctx := context.Background()
	cartService := cart.NewService(cartRepository)

	jsonFile, err := os.Open(jsonSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var items []cart.Item
	err = json.Unmarshal([]byte(jsonBytes), &items)
	if err != nil {
		panic(err)
	}

	var ids []int64
	for _, rawItem := range items {
		item, err := cartService.AddCartItem(ctx, rawItem.Name, rawItem.Price, rawItem.Manufacturer)
		if err != nil {
			panic(err)
		}
		ids = append(ids, item.ID)
	}

	return ids
}

func main() {
	var (
		dbType         = flag.String("db-type", "sqlite3", "Database Type, such as sqlite3, postgres, mysql, etc.")
		dbSource       = flag.String("db-source", "./db/my.db", "Database url connection string.")
		seedDataSource = flag.String("seed-data-source", "./db/seed.json.", "JSON Source, such as ./db/seed.json.")
	)

	flag.Parse()

	dbConn, err := driver.ConnectDB(*dbType, *dbSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	ids := SeedData(*seedDataSource, dbConn)
	fmt.Printf("ADDED %d entries.", len(ids))
}
