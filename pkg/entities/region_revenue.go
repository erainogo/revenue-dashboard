package entities

import "time"

type RegionRevenue struct {
	Region  string  `bson:"region" json:"region"`
	TotalRevenue float64 `bson:"total_revenue" json:"total_revenue"`
	TotalQuantity int     `bson:"total_quantity" json:"total_quantity"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}