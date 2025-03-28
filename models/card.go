package models

type Card struct {
	Number   string `json:"number"`
	ExpMonth string `json:"expMonth"`
	ExpYear  string `json:"expYear"`
	CVC      string `json:"cvc"`
	Updated  string `json:"updated"`
	CardType string `json:"cardType" enum:"VISA,MC,AMEX,DISC,DINERS,JCB"`

	// Optional fields
	Name               string   `json:"name,omitempty"`
	Default            bool     `json:"default,omitempty"`
	CommercialCardCode string   `json:"commercialCardCode,omitempty"`
	Address            *Address `json:"address,omitempty"`
	IsBusiness         bool     `json:"isBusiness,omitempty"`
	Created            string   `json:"created,omitempty"`
	EntityType         string   `json:"entityType,omitempty"`
	EntityId           string   `json:"entityId,omitempty"`

	// Response only fields
	IsLevel3Eligible       bool `json:"isLevel3Eligible,omitempty"`
	ZeroDollarVerification bool `json:"zeroDollarVerification,omitempty"`
}
