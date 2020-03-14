package main

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ApiKey struct {
	Model
	Key string
}

type Item struct {
	Model
	Name              string            `gorm:"unique;not null" json:"name"`
	Desc              string            `json:"desc"`
	Price             int               `json:"price"`
	ItemsTransactions []ItemTransaction `json:"items_transactions"`
}

type Transaction struct {
	Model
	Type              string            `json:"type"`
	CustomPrice       int               `json:"custom_price"`
	Cashier           string            `json:"cashier"`
	ItemsTransactions []ItemTransaction `json:"items_transactions"`
}

type ItemTransaction struct {
	Model
	Qty           int  `json:"qty"`
	ItemID        uint `json:"item_id"`
	TransactionID uint `json:"transaction_id"`
}
