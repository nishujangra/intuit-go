package client

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/models"
)

func (c *Client) CreateECheck(eCheck models.ECheck) (map[string]interface{}, error) {
	endpoint := "/quickbooks/v4/payments/echecks"

	resp, err := c.requestHelper("POST", endpoint, eCheck)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, err
	}
	return responseBody, nil
}

func (c *Client) GetECheck(eCheckID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/echecks/%s", eCheckID)

	resp, err := c.requestHelper("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, err
	}
	return responseBody, nil
}
