package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var router *mux.Router

func main() {
	fmt.Println("Starting Cozy POS backend service on port 8080!")

	InitDb(&db)
	InitRouters(&router)
	defer db.Close()

	// Populate() // If population is needed

	http.ListenAndServe(":8080", router)
}
