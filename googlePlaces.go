package main

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func saveLocationForAllDispensaries(db *gorm.DB) {
	// Instantiate the google client
	const apiKey = "AIzaSyD6_lXiJVGD79LJ8cKks1SaK4C4jDhH2u8"
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("fatal err: %s", err)
	}

	// Get dispensaries from the database
	dispensaries := getDispensaries(db)

	pretty.Printf("Dispensaries %s", dispensaries)

	// Get the dispensaries latitude/longitude
	for _, dispensary := range dispensaries {
		// TODO - pass more specific address than 'dispensary.Name'
		location := getPossibleDispensaryLatLng(dispensary.Name, client)

		// save updated dispensary to the database
		saveDispensaryLocation(dispensary, location, db)
	}
}

func getDispensaries(db *gorm.DB) []Dispensary {
	dispensaries := []Dispensary{}

	// Get all the dispensaries
	db.Find(&dispensaries)

	return dispensaries
}

func getPossibleDispensaryLatLng(possibleDispensary string, client *maps.Client) maps.LatLng {
	request := &maps.PlaceAutocompleteRequest{
		Input: possibleDispensary,
	}

	result, err := client.PlaceAutocomplete(context.Background(), request)

	if err != nil {
		log.Fatalf("fatal err: %s", err)
	}

	returnLocation := maps.LatLng{}

	for _, prediction := range result.Predictions {
		request := &maps.PlaceDetailsRequest{
			PlaceID: prediction.PlaceID,
		}

		placeDetails, err := client.PlaceDetails(context.Background(), request)

		if err != nil {
			log.Fatalf("fatal err: %s", err)
		}

		returnLocation = placeDetails.Geometry.Location

		break
	}

	return returnLocation
}

func saveDispensaryLocation(dispensary Dispensary, placeLocation maps.LatLng, db *gorm.DB) {
	db.Model(&dispensary).Updates(Dispensary{
		Latitude:  float32(placeLocation.Lat),
		Longitude: float32(placeLocation.Lng),
	})
}
