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

func (c *Client) GetCharge(chargeID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/charges/%s", chargeID)

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

func (c *Client) RefundCharge(chargeID string, refund models.Refund) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/charges/%s/refunds", chargeID)

	resp, err := c.requestHelper("POST", endpoint, refund)
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

func (c *Client) GetRefund(chargeID string, refundID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/charges/%s/refunds/%s", chargeID, refundID)

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

func (c *Client) CaptureCharge(chargeID string, capture models.Capture) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/charges/%s/capture", chargeID)

	resp, err := c.requestHelper("POST", endpoint, capture)
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

func (c *Client) VoidCharge(chargeRequestID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/quickbooks/v4/payments/txn-requests/%s/void", chargeRequestID)

	resp, err := c.requestHelper("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return data, nil
}
