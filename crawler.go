package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func CrawlSourced() {
	doc, err := goquery.NewDocument("http://sourced.tech/about/") 
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".team-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".team-item-name1").Text()
		surname := s.Find(".team-item-name2").Text()
		fmt.Printf("Worker %d: %s, %s\n", i, surname, name)
	})
}

func main() {
	CrawlSourced()
}
