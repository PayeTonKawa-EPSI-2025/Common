package models

import "time"

type Order struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	CustomerID string    `json:"customerId"`
	Products   []Product `json:"products,omitempty"` // from product service
}
