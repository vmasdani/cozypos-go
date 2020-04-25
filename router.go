package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

func InitRouters(r **mux.Router) {
	*r = mux.NewRouter()

	(*r).HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world.")
	})

	// Generate bcrypted password to put in the .env file, needs ?secret parameter

	(*r).HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query()["secret"]

		if len(key) > 0 {
			unhashedSecret := key[0]
			secret, _ := bcrypt.GenerateFromPassword([]byte(unhashedSecret), bcrypt.DefaultCost)

			secretBase64 := base64.StdEncoding.EncodeToString(secret)

			fmt.Printf("Secret to hash: %s\n", secretBase64)
			fmt.Fprintf(w, "%s", secretBase64)
		} else {
			http.Error(w, "secret paramter not found!", http.StatusInternalServerError)
			return
		}
	})

	// Old database adapter
	(*r).HandleFunc("/adapt", AdaptHandler).Methods("GET")

	// Summary
	(*r).HandleFunc("/summary", GetSummary).Methods("GET")

	// Items
	(*r).HandleFunc("/items", GetAllItems).Methods("GET")
	(*r).HandleFunc("/items/{id}", GetItem).Methods("GET")
	(*r).HandleFunc("/items-search", SearchItem).Methods("GET")
	(*r).HandleFunc("/items", PostItem).Methods("POST")
	(*r).HandleFunc("/items/stock-in", StockItemIn).Methods("POST")
	(*r).HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Transactions
	(*r).HandleFunc("/transactions", GetAllTransactions).Methods("GET")
	(*r).HandleFunc("/transactions/{id}", GetTransaction).Methods("GET")
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
	(*r).HandleFunc("/check-api-key", CheckApiKeyHandler).Methods("POST")

	// Project
	(*r).HandleFunc("/projects", GetAllProjects).Methods("GET")
	(*r).HandleFunc("/projects/{id}", GetProject).Methods("GET")
	(*r).HandleFunc("/projects/{id}/items", GetProjectItems).Methods("GET")
	(*r).HandleFunc("/projects", PostProject).Methods("POST")
	(*r).HandleFunc("/projects/{id}", DeleteProject).Methods("DELETE")

	(*r).Use(AuthorizationMiddleware)
}
