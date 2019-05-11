package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"scraper-poc/database"
	"scraper-poc/models"
	"scraper-poc/scrapers"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := database.GetDatabase()

	// Auto-create the dispensary table if it doens't exist
	db.AutoMigrate(&Dispensary{}, &Product{}, &ProductImages{})

	// Close db after the main function has finished
	defer db.Close()

	// Scrape dispensaries
	// scrapeDispensaries(db)

	// Iterate over dispensaries and update their
	// lat & lng from a place service provider
	// saveLocationForAllDispensaries(db)

	// Iterate over dispensaries and update their
	scrapers.GetAndSaveProducts(db)
}

func scrapeDispensaries(db *gorm.DB) {
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
		nameIndex := 0
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}
		addressIndex := 1
		typeIndex := 2
		activityIndex := 3

		// Loop through each table row
		tbody.ForEach("tr", func(trIndex int, tr *colly.HTMLElement) {
			// We don't care about the first 3 rows as they don't
			// contain data we need
			if trIndex >= 3 {
				dispensary := models.Dispensary{}

				// Parse each td element and construct Dispensary object
				tr.ForEach("td", func(tdIndex int, td *colly.HTMLElement) {
					if tdIndex == nameIndex {
						dispensary.Name = td.Text
					}

					if tdIndex == addressIndex {
						addressNumbersRegEx := regexp.MustCompile("[0-9-]+")
						addressNumbers := addressNumbersRegEx.FindAllString(td.Text, -1)
						streetNameRegEx := regexp.MustCompile("[A-Za-z]+")

						streetName := streetNameRegEx.FindString(td.Text)
						streetNumber := addressNumbers[0]
						zipCode := addressNumbers[1]

						dispensary.StreetName = streetName
						streetNumberParsed, _ := strconv.ParseUint(streetNumber, 0, 64)
						dispensary.StreetNumber = uint(streetNumberParsed)
						dispensary.ZipCode = zipCode
					}

					if tdIndex == typeIndex {
						dispensary.Type = td.Text
					}

					if tdIndex == activityIndex {
						if td.Text == "X" {
							dispensary.Activity = true
						}
					}

					// Create Dispensary record
					if db.NewRecord(dispensary) {
						db.Create(&dispensary)
					}
				})
			}
		})
	})

	c.Visit("https://cannabis.lacity.org/resources/authorized-retail-businesses")
}
