package models

type Charge struct {
	Token    string  `json:"token"`
	Currency string  `json:"currency"`
	Amount   string  `json:"amount"`
	Context  Context `json:"context"`

	// Optional
	Capture     string `json:"capture,omitempty" type:"boolean"`
	Description string `json:"description,omitempty"`
	CardOnFile  bool   `json:"cardOnFile,omitempty"`
	Card        Card   `json:"card,omitempty"`
}
