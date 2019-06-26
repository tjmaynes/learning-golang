package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/icrowley/fake"
	"github.com/tjmaynes/learning-golang/pkg/cart"
)

// GenerateSeedData ..
func GenerateSeedData(jsonDestination string, data []cart.Item) {
	json, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = ioutil.WriteFile(jsonDestination, json, 0644)
	if err != nil {
		panic(err)
	}
}

func generateCartItems(itemCount int, manufacturerCount int) []cart.Item {
	var manufacturers []string
	for i := 0; i < manufacturerCount; i++ {
		manufacturers = append(manufacturers, fake.Brand())
	}

	var items []cart.Item
	for i := 0; i < itemCount; i++ {
		price := cart.Decimal(int64(rand.Intn(100) + rand.Intn(100)))
		manufacturer := manufacturers[rand.Intn(4)]
		items = append(items, cart.Item{
			Name:         fake.ProductName(),
			Price:        price,
			Manufacturer: manufacturer,
		})
	}

	return items
}

func main() {
	var (
		jsonDestination   = flag.String("json-destination", "./db/seed.json", "JSON Destination, such as ./db/seed.json.")
		itemCount         = flag.Int("item-count", 50, "Number of seed items, such as 10, 20, 50, etc.")
		manufacturerCount = flag.Int("manufacturer-count", 5, "Number of unique manufacturers, such as 10, 20, 50, etc.")
	)
	flag.Parse()

	GenerateSeedData(*jsonDestination, generateCartItems(*itemCount, *manufacturerCount))
}
