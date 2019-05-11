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

// GetAndSaveProducts get and save products to db
func GetAndSaveProducts(db *gorm.DB) {
	// Loop over all the products in green rush
	// a little over 1 million entries exist
	var rowToStartFrom = 1130947
	var rowsToTraverse = 1129947

	for i := rowToStartFrom; i > rowsToTraverse; i-- {
		url := fmt.Sprintf("https://www.greenrush.com/api/v2/menu/%d?include=availability", i)
		fmt.Println("Visiting", url)
		resp, err := http.Get(url)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body", err)
		}

		greenRushProductAPIResponse, err := getProduct([]byte(body))

		if err != nil {
			fmt.Println("Error getting prodcuts from response body", err)
		}

		fmt.Println("greenRushProductAPIResponse", greenRushProductAPIResponse)
		fmt.Println("greenRushProduct", greenRushProductAPIResponse.GreenRushProduct)

		parsedProduct := parseProduct(greenRushProductAPIResponse.GreenRushProduct)

		saveProduct(parsedProduct, db)
	}
}

func getProduct(body []byte) (*GreenRushProductAPIResponse, error) {
	var s = new(GreenRushProductAPIResponse)

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error parsing GreenRushProductsAPIResponse", err)
	}

	return s, err
}

func parseProduct(product GreenRushProduct) models.Product {
	fmt.Println("Parsing product", product)
	return models.Product{
		RefID:        product.ID,
		ProductType:  product.Category.Name,
		Name:         product.Name,
		IsActive:     product.Active,
		Description:  product.Description,
		DispensaryID: product.Dispensary.ID,
	}
}

func saveProduct(product models.Product, db *gorm.DB) {
	fmt.Println("Saving product", product)
	// Create Dispensary record
	if db.NewRecord(product) {
		db.Create(&product)
	}
}
