package model

import "time"

type Item struct {
	ID int `json:"id"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
	TotalUsage int `json:"total_usage"`
	YearsInService int `json:"years_in_service"`
	NetValue float64 `json:"net_value"`
}