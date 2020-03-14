package main

import (
	"github.com/gorilla/mux"
)

type RouteHandler mux.Router

func InitRouters(r **mux.Router) {
	*r = mux.NewRouter()

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
}
