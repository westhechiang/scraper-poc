package models

import "github.com/jinzhu/gorm"

// ProductImages represents product image
type ProductImages struct {
	gorm.Model
	ProductID uint
	Link      string
}
