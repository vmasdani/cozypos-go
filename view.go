package main

type ItemTransactionView struct {
	ItemTransaction
	Transaction Transaction `json:"transaction"`
	Item        Item        `json:"item"`
}

type TransactionView struct {
	Transaction
	TotalPrice int `json:"total_price"`
}
