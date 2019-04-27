package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Dispensary represents an retailer found on
// https://cannabis.lacity.org/resources/authorized-retail-businesses
type Dispensary struct {
	gorm.Model
	Name         string
	StreetNumber uint
	StreetName   string
	ZipCode      string
	Type         string
	Activity     bool
	Latitude     float32
	Longitude    float32
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=authorized_dispensaries password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	// auto-create the dispensary table if it doens't exist
	db.AutoMigrate(&Dispensary{})

	// Close db after the main function has finished
	defer db.Close()

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
			// Original url contains query-string values that causes the google sheet to be
			// opened in another iframe, so we parse and strip the query-string
			parsedIframeURL, err := url.Parse(iframeURL)
			if err != nil {
				log.Fatal(err)
			}
			parsedIframeURL.RawQuery = ""
			modifiedURL := parsedIframeURL.String()

			detailCollector.Visit(modifiedURL)
		}
	})

	detailCollector.OnHTML("tbody", func(tbody *colly.HTMLElement) {
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			if trIndex >= 3 {
				tr.ForEach("td", func(_ int, td *colly.HTMLElement) {
					fmt.Println(td.Text)
				})
			}
		})
	})

	c.Visit("https://cannabis.lacity.org/resources/authorized-retail-businesses")
}
