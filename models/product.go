package models

import "time"

type ProductDetails struct {
	Price       float32 `json:"price" gorm:"column:price"`
	Description string  `json:"description" gorm:"column:description"`
	Color       string  `json:"color" gorm:"column:color"`
}

type Product struct {
	ID        string         `json:"id" gorm:"primaryKey;column:id"`
	Name      string         `json:"name" gorm:"column:name"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	Stock     uint           `json:"stock" gorm:"column:stock"`
	Details   ProductDetails `json:"details,omitempty" gorm:"embedded;embeddedPrefix:details_"`
}
