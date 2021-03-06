package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	fmt.Println("Starting Cozy POS backend service on port 8080!")

	InitDb(&db)
	defer db.Close()

	r := mux.NewRouter()
	InitRouters(&r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"}})
	// Debug:            true})
	// handler := cors.Default().Handler(r)
	handler := c.Handler(r)

	// Populate() // If population is needed
	http.ListenAndServe(":8080", handler)
}
