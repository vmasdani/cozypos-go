package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDb(db **gorm.DB) {
	var err error
	*db, err = gorm.Open("mysql", "valianmasdani:@/cozypos?parseTime=True")

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	(*db).SingularTable(true)

	// Migrations
	(*db).AutoMigrate(&ApiKey{})
	(*db).AutoMigrate(&Item{})
	(*db).AutoMigrate(&Transaction{})
	(*db).AutoMigrate(&ItemTransaction{})
}
