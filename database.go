package main

import "github.com/jinzhu/gorm"

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

func getDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=authorized_dispensaries password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
