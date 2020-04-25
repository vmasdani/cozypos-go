package main

import "time"

type ProjectView struct {
	ID                        uint              `json:"id"`
	UUID                      string            `json:"uuid"`
	Name                      string            `json:"name"`
	Date                      time.Time         `json:"date"`
	Transactions              []TransactionView `json:"transactions"`
	ProjectManufacturingPrice uint              `json:"project_manufacturing_price"`
	ProjectRevenue            uint              `json:"project_revenue"`
	TotalRevenue              uint              `json:"total_revenue"`
}

type TransactionView struct {
	ID                uint                  `json:"id"`
	UUID              string                `json:"uuid"`
	Type              string                `json:"type"`
	CustomPrice       int                   `json:"custom_price"`
	Cashier           string                `json:"cashier"`
	ItemsTransactions []ItemTransactionView `json:"items_transactions"`
	TotalPrice        int                   `json:"total_price"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	ProjectID         uint                  `json:"project_id"`
}

type ItemTransactionView struct {
	ID   uint   `json:"id"`
	UUID string `json:"uuid"`
	Qty  int    `json:"qty"`
	Item Item   `json:"item"`
}

type ItemView struct {
	ID                 uint          `json:"id"`
	UUID               string        `json:"uuid"`
	Name               string        `json:"name"`
	Desc               string        `json:"desc"`
	Price              int           `json:"price"`
	ManufacturingPrice int           `json:"manufacturing_price"`
	Qty                int           `json:"qty"`
	Reserved           int           `json:"reserved"`
	Sold               int           `json:"sold"`
	ItemStockIns       []ItemStockIn `json:"item_stock_ins"`
}

type ApiKeyCheck struct {
	APIKey string `json:"api_key"`
}

type Summary struct {
	TotalFund    int `json:"total_fund"`
	TotalRevenue int `json:"total_revenue"`
	TotalProfit  int `json:"total_profit"`
}
