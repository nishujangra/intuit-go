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

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetBankAccount(customerID string, accountID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/bank-accounts/%s", customerID, accountID)

	resp, err := c.requestHelper("GET", endpoint, nil)
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
