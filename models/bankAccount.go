package models

type BankAccount struct {
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType" enum:"PERSONAL_CHECKING,PERSONAL_SAVINGS,BUSINESS_CHECKING,BUSINESS_SAVINGS"`
	RoutingNumber string `json:"routingNumber" length:"9"`
	Phone         string `json:"phone"`

	// Optional fields
	BankCode  string `json:"bankCode,omitempty"`
	Country   string `json:"country,omitempty"`
	Default   bool   `json:"default,omitempty"`
	InputType string `json:"inputType,omitempty"`
}
