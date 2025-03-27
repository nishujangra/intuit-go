package client

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/models"
)

func (c *Client) CreateBankAccount(customerID string, accountData models.BankAccount) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/bank-accounts", customerID)

	resp, err := c.requestHelper("POST", endpoint, accountData)
	if err != nil {
		return nil, err
	}

	// fmt.Println(resp)

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
