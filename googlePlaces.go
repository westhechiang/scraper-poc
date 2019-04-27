package main

import (
	"context"
	"log"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

type Dispensary struct {
	Name string
}

func main() {
	// Instantiate the google client
	const apiKey = "AIzaSyD6_lXiJVGD79LJ8cKks1SaK4C4jDhH2u8"
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("fatal err: %s", err)
	}

	// Get dispensaries from the database
	dispensaries := getDispensaries()

	// Get the dispensaries latitude/longitude
	for _, dispensary := range dispensaries {
		location := getPossibleDispensaryLatLng(dispensary.Name, client)

		// save updated dispensary to the database
		saveDispensaryLocation(dispensary, location)
	}
}

func getDispensaries() []Dispensary {
	dispensaries := []Dispensary{}

	dispensary := Dispensary{
		Name: "Big Jo",
	}

	dispensaries = append(dispensaries, dispensary)

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

func saveDispensaryLocation(dispensary Dispensary, placeLocation maps.LatLng) {
	pretty.Printf("Dispensary %s", dispensary)
	pretty.Printf("Place Location: %s", placeLocation)
}
