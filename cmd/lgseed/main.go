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
		item, err := cartRepository.AddItem(ctx, &rawItem)
		if err != nil {
			panic(err)
		}
		ids = append(ids, item.ID)
	}

	return ids
}

func main() {
	var (
		dbSource   = flag.String("db-source", "mysql-user:password@/learning-golang-db", "Database url connection string.")
		dbType     = flag.String("db-type", "mysql", "Database Type, such as postgres, mysql, etc.")
		jsonSource = flag.String("json-source", "./db/seed.json.", "JSON Source, such as ./db/seed.json.")
	)

	flag.Parse()

	dbConn, err := driver.ConnectDB(*dbSource, *dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	ids := SeedData(*jsonSource, dbConn)
	fmt.Printf("ADDED %d entries.", len(ids))
}
