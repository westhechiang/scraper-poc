package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Restricting domains to visit
		colly.AllowedDomains("cannabis.lacity.org"),
	)
	c.IgnoreRobotsTxt = false
	c.Limit(&colly.LimitRule{
		DomainGlob: "https://cannabis.lacity.org/resources/authorized-retail-businesses/*",
		Delay:      2 * time.Second,
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s/n", e.Text, link)
		//  Visit link found on page
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit("https://cannabis.lacity.org/resources/authorized-retail-businesses")
}
