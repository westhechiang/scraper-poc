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
	LegalStatus  string
}

type Product struct {
	gorm.Model
	RefID       uint
	ProductType string
	Name        string
	Status      string
	// Category    string // Cat
	ProductTypeID uint
	Description   string
	DispensaryID  uint
}

type ProductImages struct {
	gorm.Model
	ProductID uint
	Link      string
}

func getDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=authorized_dispensaries password=postgres sslmode=disable")
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres-dev dbname=dev password=scraper1 sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
