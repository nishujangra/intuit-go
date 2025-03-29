package models

type Charge struct {
	Token    string  `json:"token"`
	Currency string  `json:"currency"`
	Amount   string  `json:"amount"`
	Context  Context `json:"context"`

	// Optional
	Capture     bool   `json:"capture,omitempty"`
	Description string `json:"description,omitempty"`
	CardOnFile  bool   `json:"cardOnFile,omitempty"`
	Card        Card   `json:"card,omitempty"`
}

type Context struct {
	Mobile      string `json:"mobile,omitempty" type:"boolean"`
	IsEcommerce string `json:"isEcommerce,omitempty" type:"boolean"`
}
