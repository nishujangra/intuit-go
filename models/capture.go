package models

type Capture struct {
	Amount      string   `json:"amount"`
	Description string   `json:"description"`
	Context     *Context `json:"context"`
}
