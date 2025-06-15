package entities

import "time"

type ProductPurchaseSummary struct {
	ProductID       string    `bson:"product_id" json:"product_id"`
	ProductName     string    `bson:"product_name" json:"product_name"`
	PurchaseCount   int       `bson:"purchase_count" json:"purchase_count"`
	StockQuantity   int       `bson:"stock_quantity" json:"stock_quantity"`
	UpdatedAt       time.Time `bson:"updated_at" json:"updated_at"`
}
