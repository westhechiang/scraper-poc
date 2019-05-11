package database

import "github.com/jinzhu/gorm"

func GetDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=authorized_dispensaries password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
