package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

const QUERY = "?sz=1000&format=ajax"

type Item struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Position    int    `json:"position"`
	List        string `json:"list"`
	Brand       string `json:"brand"`
	Variant     string `json:"variant"`
	Quantity    int    `json:"quantity"`
	ProductCode string `json:"productCode"`
}

func main() {
	var url string

	flag.StringVar(&url, "url", "", "URL")
	flag.Parse()

	if url == "" {
		log.Fatal("--url must be specified")
	}

	db, err := sql.Open("postgres", "user=striketime dbname=striketime sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loading %s... ", url)

	doc, err := goquery.NewDocument(url + QUERY)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")

	doc.Find(".product-tile").Each(func(i int, s *goquery.Selection) {
		item := &Item{}
		itemJSON, _ := s.Attr("data-line-item")
		json.Unmarshal([]byte(itemJSON), item)

		image, _ := s.Find("img").First().Attr("src")
		onSale := s.Find(".price-sales").Size() > 0

		if _, err := db.Exec("INSERT INTO products (id, name, brand, category, variant, code, image) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO UPDATE SET (id, name, brand, category, variant, code, image) = ($1, $2, $3, $4, $5, $6, $7)", item.ID, item.Name, item.Brand, item.Category, item.Variant, item.ProductCode, image); err != nil {
			log.Println(err)
			return
		}

		var existingPrice float64
		currentPrice, _ := strconv.ParseFloat(item.Price, 64)

		err := db.QueryRow("SELECT price FROM product_prices WHERE product_id = $1", item.ID).Scan(&existingPrice)
		if err != nil || existingPrice != currentPrice {
			_, err = db.Exec("INSERT INTO product_prices (product_id, price, sale) VALUES ($1, $2, $3)", item.ID, currentPrice, onSale)
			if err != nil {
				log.Println(err)
			}
		}
	})
}
