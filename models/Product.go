package models

import "github.com/jinzhu/gorm"

// Product represents product image
type Product struct {
	gorm.Model
	RefID       uint
	ProductType string
	Name        string
	IsActive    bool
	// Category    string // Cat
	ProductTypeID uint
	Description   string
	DispensaryID  uint
}
