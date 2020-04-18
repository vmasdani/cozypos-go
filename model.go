package main

import (
	"time"
)

// LoginInfo : Login info request struct
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Model : Essential GORM model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// ApiKey : table for storing API keys
type APIKey struct {
	Model
	APIKey string `json:"api_key"`
}

// Item : table for storing items/inventory list
type Item struct {
	Model
	Name               string            `gorm:"unique;not null" json:"name"`
	Desc               string            `json:"desc"`
	Price              int               `json:"price"`
	ManufacturingPrice int               `json:"manufacturing_price"`
	ItemsTransactions  []ItemTransaction `json:"items_transactions"`
}

// Transaction : table for storing transactions
type Transaction struct {
	Model
	Type              string            `json:"type"`
	CustomPrice       int               `json:"custom_price"`
	Cashier           string            `json:"cashier"`
	ItemsTransactions []ItemTransaction `json:"items_transactions"`
	ProjectID         uint              `json:"project_id"`
}

// ItemTransaction : table for storing ItemTransaction
type ItemTransaction struct {
	Model
	Qty           int  `json:"qty"`
	ItemID        uint `json:"item_id"`
	TransactionID uint `json:"transaction_id"`
}

// Project : table for storing projects
type Project struct {
	Model
	Name         string        `json:"name"`
	Date         time.Time     `json:"date"`
	Transactions []Transaction `json:"transactions"`
}
