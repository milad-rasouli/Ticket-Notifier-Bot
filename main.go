package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://mrbilit.com/buses/mashhad-tehran?departureDate=1403-03-09`),
		chromedp.WaitVisible(`div.trip-card-wrapper`, chromedp.ByQuery),
		chromedp.OuterHTML(`div.trip-card-wrapper`, &htmlContent, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(htmlContent)
}
