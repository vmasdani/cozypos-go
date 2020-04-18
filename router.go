package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouters(r **mux.Router) {
	*r = mux.NewRouter()

	(*r).HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world.")
	})

	// Items
	(*r).HandleFunc("/items", GetAllItems).Methods("GET")
	(*r).HandleFunc("/items", PostItem).Methods("POST")
	(*r).HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Transactions
	(*r).HandleFunc("/transactions", GetAllTransactions).Methods("GET")
	(*r).HandleFunc("/transactions", PostTransaction).Methods("POST")
	(*r).HandleFunc("/transactions/{id}", DeleteTransaction).Methods("DELETE")

	// ItemsTransactions
	(*r).HandleFunc("/items-transactions", GetAllItemsTransactions).Methods("GET")
	(*r).HandleFunc("/items-transactions", PostItemTransaction).Methods("POST")
	(*r).HandleFunc("/items-transactions/{id}", DeleteItemTransaction).Methods("DELETE")

	// API key
	(*r).HandleFunc("/api-keys", GetAllApiKeys).Methods("GET")
	(*r).HandleFunc("/api-keys", PostApiKey).Methods("POST")
	(*r).HandleFunc("/api-keys/{id}", DeleteApiKey).Methods("DELETE")

	// Login
	(*r).HandleFunc("/login", LoginHandler).Methods("POST")

	// Project
	(*r).HandleFunc("/projects", GetAllProjects).Methods("GET")
	(*r).HandleFunc("/projects/{id}", GetProject).Methods("GET")
	(*r).HandleFunc("/projects", PostProject).Methods("POST")
	(*r).HandleFunc("/projects/{id}", DeleteProject).Methods("DELETE")

	(*r).Use(AuthorizationMiddleware)
}
