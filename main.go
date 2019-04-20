package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = false
	c.Limit(&colly.LimitRule{
		// Set a delay between requests
		Delay: 2 * time.Second,
		// Set additional random delay
		RandomDelay: 1 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	detailCollector := c.Clone()

	// Finds all iframes (https://cannabis.lacity.org/resources/authorized-retail-businesses uses
	// Google Sheets embedded via iframe to host list of legal retailers)
	c.OnHTML("iframe", func(e *colly.HTMLElement) {
		iframeURL := e.Attr("src")
		// Only visit iframe src if it's google docs, assuming 'google' in url for now
		if strings.Contains(iframeURL, "google") {
			parsedIframeURL, err := url.Parse(iframeURL)
			if err != nil {
				log.Fatal(err)
			}
			parsedIframeURL.RawQuery = ""
			modifiedURL := parsedIframeURL.String()

			detailCollector.Visit(modifiedURL)
		}
	})

	detailCollector.OnHTML("tr", func(r *colly.HTMLElement) {
		fmt.Println("2nd iframe")
		fmt.Println(r.Text)
	})

	c.Visit("https://cannabis.lacity.org/resources/authorized-retail-businesses")
}
