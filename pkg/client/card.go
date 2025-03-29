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

func (c *Client) GetCard(customerID, cardId string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/cards/%s", customerID, cardId)

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

func (c *Client) GetCards(customerID string) ([]map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/cards", customerID)

	resp, err := c.requestHelper("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DeleteCard(customerID, cardId string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/customers/%s/cards/%s", customerID, cardId)

	resp, err := c.requestHelper("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Card deleted successfully",
		"code":    resp.StatusCode,
	}, nil
}
