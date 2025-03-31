package models

type ECheck struct {
	Amount       string       `json:"amount"`
	PayementMode string       `json:"paymentMode" enum:"WEB"`
	BankAccount  *BankAccount `json:"bankAccount"`

	Description string   `json:"description"`
	CheckNumber string   `json:"checkNumber"`
	Context     *Context `json:"context"`
}
