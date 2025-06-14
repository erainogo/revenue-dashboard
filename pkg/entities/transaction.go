package entities

import "time"

type Product struct {
	ID            string `bson:"id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Category      string `bson:"category" json:"category"`
	StockQuantity int    `bson:"stock_quantity" json:"stock_quantity"`
}

type Transaction struct {
	TransactionID   string      `bson:"transaction_id" json:"transaction_id"`
	TransactionDate time.Time   `bson:"transaction_date" json:"transaction_date"`
	UserID          string      `bson:"user_id" json:"user_id"`
	Country         string      `bson:"country" json:"country"`
	Region          string      `bson:"region" json:"region"`
	Product         Product     `bson:"product" json:"product"`
	Price           float64     `bson:"price" json:"price"`
	Quantity        int         `bson:"quantity" json:"quantity"`
	TotalPrice      float64     `bson:"total_price" json:"total_price"`
	AddedDate       time.Time    `bson:"added_date" json:"added_date"`
}

type SummaryKey struct {
	Country     string
	ProductName string
}