package entities

import "time"

type MonthlySales struct {
	Year          int  `bson:"year" json:"year"`
	Month         string  `bson:"month" json:"month"`
	TotalQuantity int     `bson:"total_quantity" json:"total_quantity"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}
