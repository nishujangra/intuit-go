package client

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/models"
)

func (c *Client) CreateCharge(charge models.Charge) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/charges")

	resp, err := c.requestHelper("POST", endpoint, charge)
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
