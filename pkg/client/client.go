package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nishujangra/intuit-go/config"
)

type Client struct {
	BaseURL     string
	AccessToken string
	HTTPClient  *http.Client
}

func NewClient(accessToken string) *Client {
	return &Client{
		BaseURL:     config.GetPaymentsBaseURL(),
		AccessToken: accessToken,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) requestHelper(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	if method == "POST" {
		req.Header.Set("request-id", uuid.New().String())
	}

	return c.HTTPClient.Do(req)
}
