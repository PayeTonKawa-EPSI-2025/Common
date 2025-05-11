package models

import "time"

type Address struct {
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Company struct {
	CompanyName string `json:"companyName"`
}

type Customer struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`

	Address *Address `json:"address,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
	Company *Company `json:"company,omitempty"`

	Orders []Order `json:"orders,omitempty"` // from order service
}
