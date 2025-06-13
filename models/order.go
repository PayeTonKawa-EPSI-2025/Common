package models

import "time"

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	CustomerID string    `json:"customerId" gorm:"column:customer_id"`
	Products   []Product `json:"products,omitempty" gorm:"-"`
}
