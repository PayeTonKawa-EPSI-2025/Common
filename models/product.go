package models

import "time"

type ProductDetails struct {
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
}

type Product struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"createdAt"`
	Stock     uint            `json:"stock"`
	Details   *ProductDetails `json:"details,omitempty"`
}
