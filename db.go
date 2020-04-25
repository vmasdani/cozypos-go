package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
)

func InitDb(db **gorm.DB) {
	var dbErr error
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		fmt.Println(dotenvErr)
		panic("Failed loading dotenv.")
	}

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbName)

	*db, dbErr = gorm.Open("mysql", dbUrl)

	if dbErr != nil {
		fmt.Println(dbErr)
		panic("Failed to connect to database!")
	}

	(*db).SingularTable(true)

	// Migrations
	(*db).AutoMigrate(&APIKey{})
	(*db).AutoMigrate(&Item{})
	(*db).AutoMigrate(&Transaction{})
	(*db).AutoMigrate(&ItemTransaction{})
	(*db).AutoMigrate(&Project{})
	(*db).AutoMigrate(&ItemProject{})
	(*db).AutoMigrate(&ItemStockIn{})
}
