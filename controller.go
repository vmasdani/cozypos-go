package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

// Item
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allItems []Item
	db.Order("id desc").Preload("ItemsTransactions").Preload("ItemStockIns").Find(&allItems)

	itemViews := []ItemView{}

	for _, item := range allItems {
		totalQty := 0
		totalSold := 0
		totalReserved := 0

		// Count total sold
		for _, itemTransaction := range item.ItemsTransactions {
			totalSold += itemTransaction.Qty
		}

		// Count total qty
		for _, itemStockIn := range item.ItemStockIns {
			totalQty += itemStockIn.Qty
		}

		// Count total reserved
		for _, itemProject := range item.ItemProjects {
			totalReserved += itemProject.Qty
		}

		newItemView := ItemView{
			ID:                 item.ID,
			Name:               item.Name,
			Desc:               item.Desc,
			Price:              item.Price,
			ManufacturingPrice: item.ManufacturingPrice,
			Qty:                totalQty,
			Reserved:           totalReserved,
			Sold:               totalSold}

		newItemViews := append(itemViews, newItemView)
		itemViews = newItemViews
	}

	json.NewEncoder(w).Encode(itemViews)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var foundItem Item
	db.Preload("ItemsTransactions").Preload("ItemStockIns").First(&foundItem, id)

	totalQty := 0

	// Count item in
	for _, itemStockIn := range foundItem.ItemStockIns {
		totalQty += itemStockIn.Qty
	}

	// Count item sold
	for _, itemTransaction := range foundItem.ItemsTransactions {
		totalQty -= itemTransaction.Qty
	}

	itemView := ItemView{
		ID:                 foundItem.ID,
		Name:               foundItem.Name,
		Desc:               foundItem.Desc,
		Price:              foundItem.Price,
		ManufacturingPrice: foundItem.ManufacturingPrice,
		Qty:                totalQty}

	json.NewEncoder(w).Encode(itemView)
}

func SearchItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemName := r.URL.Query()["name"]

	if len(itemName) > 0 {
		itemLike := fmt.Sprintf("%%%s%%", itemName[0])
		var foundItems []Item
		db.Limit(10).Where("name LIKE ?", itemLike).Find(&foundItems)

		json.NewEncoder(w).Encode(foundItems)
	} else {
		http.Error(w, "Parameter name does not exist.", http.StatusBadRequest)
	}

	// json.NewEncoder(w).Encode()
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	insertion := db.Save(&item)

	if insertion.Error != nil {
		http.Error(w, "Failed to create new item.", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(item)
}

func StockItemIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Body)

	var itemStockIn ItemStockIn
	json.NewDecoder(r.Body).Decode(&itemStockIn)

	insertion := db.Save(&itemStockIn)

	if insertion.Error != nil {
		http.Error(w, "Error stocking in.", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(itemStockIn)
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
	// db.Preload("ItemsTransactions").Find(&transactions)
	db.Find(&transactions)

	// transactionViews := []TransactionView{}

	// for _, transaction := range transactions {

	// 	// Sum all accumulated item transactions
	// 	totalPrice := 0

	// 	for _, itemTransaction := range transaction.ItemsTransactions {
	// 		var foundItem Item
	// 		db.Where("id = ?", itemTransaction.ItemID).First(&foundItem)

	// 		totalPrice += itemTransaction.Qty * foundItem.Price
	// 	}

	// 	transactionViews = append(transactionViews, TransactionView{
	// 		ID:          transaction.ID,
	// 		Type:        transaction.Type,
	// 		CustomPrice: transaction.CustomPrice,
	// 		Cashier:     transaction.Cashier,
	// 		TotalPrice:  totalPrice})
	// }

	// Finally, serialize the TransactionView
	json.NewEncoder(w).Encode(transactions)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	totalPrice := 0

	var transaction Transaction
	if db.First(&transaction, id).RecordNotFound() {
		http.Error(w, "Transaction not found", http.StatusInternalServerError)
	}

	var itemsTransactions []ItemTransaction
	db.Where("transaction_id = ?", transaction.ID).Find(&itemsTransactions)

	itemTransactionViews := []ItemTransactionView{}

	for _, itemTransaction := range itemsTransactions {
		var foundItem Item
		db.First(&foundItem, itemTransaction.ItemID)

		itemTransactionView := ItemTransactionView{
			ID:   itemTransaction.ID,
			Qty:  itemTransaction.Qty,
			Item: foundItem}

		newItemTransactionViews := append(itemTransactionViews, itemTransactionView)
		itemTransactionViews = newItemTransactionViews

		// Update total price:
		totalPrice += itemTransaction.Qty * foundItem.Price
	}

	transactionView := TransactionView{
		ID:                transaction.ID,
		Type:              transaction.Type,
		CustomPrice:       transaction.CustomPrice,
		Cashier:           transaction.Cashier,
		ItemsTransactions: itemTransactionViews,
		TotalPrice:        totalPrice,
		CreatedAt:         transaction.CreatedAt,
		UpdatedAt:         transaction.UpdatedAt,
		ProjectID:         transaction.ProjectID}

	fmt.Println(transactionView)

	json.NewEncoder(w).Encode(transactionView)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Username: %s\n", r.Header.Get("Username"))

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
			ID:   itemTransaction.ID,
			Qty:  itemTransaction.Qty,
			Item: foundItem})
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

// API key
func GetAllApiKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var apiKeys []APIKey

	db.Find(&apiKeys)

	json.NewEncoder(w).Encode(apiKeys)
}

func PostApiKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var apiKey APIKey
	json.NewDecoder(r.Body).Decode(&apiKey)

	// Create new API key
	db.Save(&apiKey)

	json.NewEncoder(w).Encode(apiKey)
}

func DeleteApiKey(w http.ResponseWriter, r *http.Request) {
	var apiKeyId = mux.Vars(r)["id"]

	var apiKey APIKey
	db.Where("id = ?", apiKeyId).First(&apiKey)

	db.Delete(apiKey)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginInfo LoginInfo
	json.NewDecoder(r.Body).Decode(&loginInfo)

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error getting env file")
		http.Error(w, "Error getting env file.", http.StatusInternalServerError)
		return
	}
	secretCodeBase64 := os.Getenv("SECRET")
	secretCode, _ := base64.StdEncoding.DecodeString(secretCodeBase64)
	secretCodeBytes := []byte(secretCode)
	passwordBytes := []byte(loginInfo.Password)

	loginErr := bcrypt.CompareHashAndPassword(secretCodeBytes, passwordBytes)
	secretHash, generateSecretError := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if loginErr != nil {
		fmt.Println("Password do not match!")
		http.Error(w, "Password incorrect!", http.StatusUnauthorized)
		return
	}

	if generateSecretError == nil {
		fmt.Println(string(secretHash))
	}

	// Generate api key
	timestamp := time.Now().Unix()
	randNum := rand.Float64()

	randToStr := fmt.Sprintf("%d%e", timestamp, randNum)

	unhashedApiKey := []byte(randToStr)

	keyBytes, err := bcrypt.GenerateFromPassword(unhashedApiKey, bcrypt.DefaultCost)
	keyBase64 := base64.StdEncoding.EncodeToString(keyBytes)

	usernameBytes := []byte(loginInfo.Username)
	usernameBase64 := base64.StdEncoding.EncodeToString(usernameBytes)

	apiKey := fmt.Sprintf("%s:%s", usernameBase64, keyBase64)

	fmt.Println(usernameBase64)
	fmt.Println(keyBase64)

	// bcrypt.GeneratePassword([]byte())

	fmt.Printf("Secret code: %s\n", secretCode)
	fmt.Printf("Api key: %s\n", apiKey)

	fmt.Fprintf(w, "%s", apiKey)
}

// Project
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var projects []Project
	// db.Preload("Transactions").Find(&projects)
	db.Find(&projects)

	json.NewEncoder(w).Encode(projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	projectManufacturingPrice := 0
	totalRevenue := 0
	projectRevenue := 0

	fmt.Println(projectManufacturingPrice)
	fmt.Println(totalRevenue)
	fmt.Println(projectRevenue)

	var project Project
	db.Where("id = ?", id).Preload("Transactions").First(&project)

	transactionViews := []TransactionView{}

	// Loop transactions
	for _, transaction := range project.Transactions {
		var foundTransaction Transaction
		db.Preload("ItemsTransactions").First(&foundTransaction, transaction.ID)

		// Count total price
		totalPrice := 0
		manufacturingPrice := 0

		itemTransactionViews := []ItemTransactionView{}

		for _, itemTransaction := range foundTransaction.ItemsTransactions {
			var foundItem Item
			db.First(&foundItem, itemTransaction.ItemID)

			totalPrice += itemTransaction.Qty * foundItem.Price

			itemTransactionView := ItemTransactionView{
				ID:   itemTransaction.ID,
				Qty:  itemTransaction.Qty,
				Item: foundItem}

			newItemTransactionViews := append(itemTransactionViews, itemTransactionView)
			itemTransactionViews = newItemTransactionViews

			// Update manufacturing price
			manufacturingPrice += itemTransaction.Qty * foundItem.ManufacturingPrice
		}

		transactionView := TransactionView{
			ID:          transaction.ID,
			Type:        transaction.Type,
			CustomPrice: transaction.CustomPrice,
			Cashier:     transaction.Cashier,
			TotalPrice:  totalPrice}
		// ItemsTransactions: itemTransactionViews} // Uncomment this to debug items transactions in project

		newTransactionViews := append(transactionViews, transactionView)
		transactionViews = newTransactionViews

		// Update total revenue
		switch transaction.Type {
		case "sell":
			projectRevenue += totalPrice
		case "stock_in":
			projectManufacturingPrice += manufacturingPrice
		case "auction":
			projectRevenue += transaction.CustomPrice
		default:
			fmt.Printf("Transaction ID: %d does not belong into any transaction type.\n", transaction.ID)
		}
	}

	projectView := ProjectView{
		ID:                        project.ID,
		Name:                      project.Name,
		Date:                      project.Date,
		TotalRevenue:              uint(totalRevenue),
		ProjectRevenue:            uint(projectRevenue),
		ProjectManufacturingPrice: uint(projectManufacturingPrice),
		Transactions:              transactionViews}

	fmt.Println("Project:")
	fmt.Println(projectView)

	// fmt.Println("Transactions:")
	// for _, transaction := range project.Transactions {
	// 	fmt.Println(transaction)
	// }

	json.NewEncoder(w).Encode(&projectView)
}

func GetProjectItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var foundProject Project
	db.Preload("ItemProjects").First(&foundProject, id)

	json.NewEncoder(w).Encode(foundProject)
}

func PostProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, "Error parsing Request!", http.StatusInternalServerError)
		return
	}

	db.Save(&project)
	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var project Project
	db.First(&project, id)

	db.Delete(&project)
}
