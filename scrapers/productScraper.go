package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"scraper-poc/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// https://www.greenrush.com/api/v2/menu/1 <<<< Min product ID
// https://www.greenrush.com/api/v2/menu/1130933 <<<< Max product ID

func GetAndSaveProducts(db *gorm.DB) {
	// Loop over all the products in green rush
	// a little over 1 million entries exist
	for i := 1; i < 2; i++ {
		url := fmt.Sprintf("https://www.greenrush.com/api/v2/menu/%d?include=availability", i)
		fmt.Println("Visiting", url)
		resp, err := http.Get(url)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body", err)
		}

		products, err := getProducts([]byte(body))
		if err != nil {
			fmt.Println("Error getting prodcuts from response body", err)
		}
	}
}

// GreenRushProductsAPIResponsee
type GreenRushProductsAPIResponse struct {
	Products []models.Product `json:"data`
}

func getProducts(ody []byte) (*GreenRushProductsAPIResponse, error) {
	var s = new(GreenRushProductsAPIResponse)

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error parsing GreenRushProductsAPIResponse", err)
	}

	return s, err
}
