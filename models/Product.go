package models

import "github.com/jinzhu/gorm"

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
