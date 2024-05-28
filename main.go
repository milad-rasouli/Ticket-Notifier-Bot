package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	var (
		err error
		url = "https://mrbilit.com/buses/mashhad-tehran?departureDate=1403-03-09"
	)
	c := colly.NewCollector(colly.AllowedDomains("mrbilit.com", "www.mrbilit.com"))

	// Find and visit all links
	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Println("OnHTML ", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("%w", err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Request: %s", r.URL)
	})
	
	err = c.Visit(url)
	if err != nil {
		log.Println(err)
	}
}
