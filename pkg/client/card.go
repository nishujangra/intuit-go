package client

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/models"
)

func (c *Client) CreateCard(customerID string, cardData models.Card) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/cards", customerID)

	resp, err := c.requestHelper("POST", endpoint, cardData)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
