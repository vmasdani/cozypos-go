package main

import "time"

type ProjectView struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	Date         time.Time         `json:"Date"`
	Transactions []TransactionView `json:"transactions"`
}

type TransactionView struct {
	ID                uint                  `json:"id"`
	Type              string                `json:"type"`
	CustomPrice       int                   `json:"custom_price"`
	Cashier           string                `json:"cashier"`
	ItemsTransactions []ItemTransactionView `json:"items_transactions"`
	TotalPrice        int                   `json:"total_price"`
}

type ItemTransactionView struct {
	ID   uint `json:"id"`
	Qty  int  `json:"qty"`
	Item Item `json:"item"`
}
