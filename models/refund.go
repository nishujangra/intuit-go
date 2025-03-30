package models

type Refund struct {
	Amount      string   `json:"amount"`
	Description string   `json:"description,omitempty"`
	Context     *Context `json:"context,omitempty"`
}
