package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nishujangra/intuit-go/config"
)

type Client struct {
	BaseURL     string
	AccessToken string
	HTTPClient  *http.Client
	Logger      *log.Logger
}

func NewClient(accessToken string) *Client {

	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		BaseURL:     config.GetPaymentsBaseURL(),
		AccessToken: accessToken,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Logger: log.New(logFile, "INFO: ", log.Ldate|log.Ltime),
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
	req.Header.Set("Accept", "application/json")

	requestID := uuid.New().String()

	if method == "POST" {
		req.Header.Set("request-id", requestID)
	}

	c.Logger.Printf(
		"Response Status: %s, Request ID: %s, URL: %s, Method: %s",
		"Requesting",
		requestID,
		url,
		method,
	)

	return c.HTTPClient.Do(req)
}
