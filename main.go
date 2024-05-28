package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		fmt.Println("OnHTML ", e.Attr("href"))
		fmt.Println("OnHTML name: ", e.Name)
		//e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://safar724.com/bus/mashhad-gonabad?date=1403-03-08")
}
