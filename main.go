package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

func main() {
	var url string

	flag.StringVar(&url, "url", "", "URL to fetch")
	flag.Parse()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// extract the metadata
	idString := doc.Find(".product-id").First().Text()
	id, _ := strconv.Atoi(strings.Trim(idString, " \n"))

	priceString, _ := doc.Find(".pv-price").First().Attr("data-price")
	price, _ := strconv.ParseFloat(strings.Trim(priceString, " \n"), 32)

	imageSrc, _ := doc.Find(".primary-image").First().Attr("src")

	productName := doc.Find(".product-name").First().Text()

	availability := doc.Find("#availability-value").First().Text()
	availability = strings.Trim(availability, " \n")

	offersEnd := doc.Find(".offers-end").First().Text()
	offersEnd = strings.TrimPrefix(strings.Trim(offersEnd, " \n"), "Offer Ends: ")

	promotionCallout := doc.Find(".promotion-callout").First().Text()
	promotionCallout = strings.Trim(promotionCallout, " \n")

	description, _ := doc.Find(".description-text").First().Html()
	description = strings.Trim(description, " \n")

	fmt.Println("ID:", id)
	fmt.Println("Name:", productName)
	fmt.Println("Availability:", availability)
	fmt.Printf("Price: $%.2f\n", price)
	fmt.Println("Promotion:", promotionCallout)
	fmt.Println("Offers End:", offersEnd)
	fmt.Println("Image:", imageSrc)
	fmt.Println("Description:", description)
}
