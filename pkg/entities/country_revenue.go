package entities

import "time"

type CountryLevelRevenue struct {
	Country          string     `bson:"country" json:"country"`
	ProductName      string     `bson:"product_name" json:"product_name"`
	TotalRevenue     float64    `bson:"total_revenue" json:"total_revenue"`
	TransactionCount int32      `bson:"transaction_count" json:"transaction_count"`
	UpdatedAt        time.Time  `bson:"updated_at" json:"updated_at"`
}

