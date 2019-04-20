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

	detailCollector := c.Clone()

	c.OnHTML("iframe", func(e *colly.HTMLElement) {
		iframeURL := e.Attr("src")
		// Print link
		// fmt.Printf("Link found: %q -> %s/n", e.Text, link)
		//  Visit link found on page
		if strings.Contains(iframeURL, "google") {
			fmt.Println("1st iframe", iframeURL)
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

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(r.Body)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://cannabis.lacity.org/resources/authorized-retail-businesses")
}
