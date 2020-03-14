package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Item
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var allItems []Item
	db.Preload("ItemsTransactions").Find(&allItems)

	fmt.Println(allItems)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allItems)
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	db.Save(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := mux.Vars(r)["id"]

	var item Item
	db.Where("id = ?", itemId).First(&item)

	db.Delete(&item)
}

// Transaction
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transactions []Transaction
	db.Preload("ItemsTransactions").Find(&transactions)

	transactionViews := []TransactionView{}

	for _, transaction := range transactions {

		// Sum all accumulated item transactions
		totalPrice := 0

		for _, itemTransaction := range transaction.ItemsTransactions {
			var foundItem Item
			db.Where("id = ?", itemTransaction.ItemID).First(&foundItem)

			totalPrice += itemTransaction.Qty * foundItem.Price
		}

		transactionViews = append(transactionViews, TransactionView{
			Transaction: transaction,
			TotalPrice:  totalPrice})
	}

	// Finally, serialize the TransactionView
	json.NewEncoder(w).Encode(transactionViews)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transaction Transaction
	json.NewDecoder(r.Body).Decode(&transaction)

	fmt.Println("Transaction to save", transaction)

	db.Save(&transaction)

	json.NewEncoder(w).Encode(transaction)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId := mux.Vars(r)["id"]

	var transaction Transaction
	db.Where("id = ?", transactionId).First(&transaction)

	db.Delete(&transaction)
}

// ItemTransaction
func GetAllItemsTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var itemsTransactions []ItemTransaction
	db.Find(&itemsTransactions)

	itemsTransactionsView := []ItemTransactionView{}

	for _, itemTransaction := range itemsTransactions {
		var foundItem Item
		var foundTransaction Transaction

		db.Where("id = ?", itemTransaction.ItemID).First(&foundItem)
		db.Where("id = ?", itemTransaction.TransactionID).First(&foundTransaction)

		itemsTransactionsView = append(itemsTransactionsView, ItemTransactionView{
			ItemTransaction: itemTransaction,
			Transaction:     foundTransaction,
			Item:            foundItem})
	}

	json.NewEncoder(w).Encode(itemsTransactionsView)
}

func PostItemTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Body)

	var itemTransaction ItemTransaction
	json.NewDecoder(r.Body).Decode(&itemTransaction)

	fmt.Println("item transaction to insert", itemTransaction)

	db.Save(&itemTransaction)

	json.NewEncoder(w).Encode(itemTransaction)
}

func DeleteItemTransaction(w http.ResponseWriter, r *http.Request) {
	itemTransactionId := mux.Vars(r)["id"]

	var itemTransaction ItemTransaction
	db.Where("id = ?", itemTransactionId).First(&itemTransaction)

	db.Delete(&itemTransaction)
}
