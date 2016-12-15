package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func main() {
	doc, err := goquery.NewDocument("https://www.thewarehouse.co.nz")
	if err != nil {
		log.Fatal(err)
	}

	categories := make([]string, 0)

	doc.Find(".menu-container-level-3 a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		categories = appendIfMissing(categories, href)
	})

	for _, url := range categories {
		fmt.Println(url)
	}
}
