package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandler mux.Router

func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We are accepting.")

		for header := range r.Header {
			fmt.Printf("%s: %s\n", header, r.Header.Get(header))
		}

		fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))
		fmt.Printf("Connection: %s\n", r.Header.Get("Connection"))

		next.ServeHTTP(w, r)
	})
}

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

	(*r).Use(authorizationMiddleware)
}
