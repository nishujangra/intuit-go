package models

type Address struct {
	PostalCode string `json:"postalCode"`
	Region     string `json:"region"`

	// Optional fields
	StreetAddress string `json:"streetAddress,omitempty"`
	City          string `json:"city,omitempty"`
	Country       string `json:"country,omitempty"`
}
