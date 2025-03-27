package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/nishujangra/intuit-go/config"
)

type AuthClient struct {
	CLIENT_ID       string
	CLIENT_SECRET   string
	TOKEN_END_POINT string
	HTTPClient      *http.Client
}

func NewAuthClient(clientID, clientSecret, tokenURL string) *AuthClient {
	return &AuthClient{
		CLIENT_ID:       clientID,
		CLIENT_SECRET:   clientSecret,
		TOKEN_END_POINT: "/oauth2/v1/tokens/bearer",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ac *AuthClient) GetToken(authCode string) (string, error) {
	baseUrl := config.GetPaymentsAuthURL()
	urlStr := baseUrl + ac.TOKEN_END_POINT

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", "https://developer.intuit.com/v2/OAuth2Playground/RedirectUrl")

	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", errors.New("Error creating request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	authHeader := base64.StdEncoding.EncodeToString([]byte(ac.CLIENT_ID + ":" + ac.CLIENT_SECRET))
	req.Header.Set("Authorization", "Basic "+authHeader)

	resp, err := ac.HTTPClient.Do(req)
	if err != nil {
		return "", errors.New("Error sending request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Error response")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", errors.New("Error decoding response")
	}

	return result["access_token"].(string), nil
}
