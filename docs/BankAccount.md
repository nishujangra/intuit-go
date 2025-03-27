# Bank Account API - QuickBooks Intuit SDK

This document provides details on how to create a **Bank Account** using the Intuit API in our SDK.

## 🌐 API Endpoint

**POST** `/quickbooks/v4/customers/{customer-id}/bank-accounts`

- **Production URL:** `https://api.intuit.com`
- **Sandbox URL:** `https://sandbox.api.intuit.com`

## 🔑 Headers
| Header       | Value Example             | Description |
|-------------|--------------------------|-------------|
| Authorization | `Bearer {ACCESS_TOKEN}`  | Your OAuth2 Access Token |
| Content-Type | `application/json`       | API request format |
| request-Id  | `UUID`                    | Unique request identifier |

## 📥 Request Body
| Parameter       | Required | Type   | Example Value        | Description |
|----------------|----------|--------|----------------------|-------------|
| `name`         | ✅       | String | `"Tony Stark"`         | Name of account holder |
| `accountNumber`| ✅       | String | `"123456789012"`     | 4-17 digit account number |
| `accountType`  | ✅       | Enum   | `"PERSONAL_CHECKING"` | Type of account (Checking/Savings) |
| `routingNumber`| ✅       | String | `"021000021"`        | 9-digit routing number |
| `phone`        | ✅       | String | `"9876543210"`       | 10-digit phone number |
| `default`      | ❌       | Boolean| `true`               | Whether this is the default account |
| `country`      | ❌       | String | `"US"`               | ISO country code |
| `bankCode`     | ❌       | String | `"021000021"`        | Bank code (varies by country) |
| `inputType`    | ❌       | Enum   | `"KEYED"`            | Method of input |


## 🛠 Example Request
```json
{
  "name": "John Doe",
  "accountNumber": "123456789012",
  "accountType": "PERSONAL_CHECKING",
  "routingNumber": "021000021",
  "phone": "9876543210",
  "country": "US",
  "default": true,
  "bankCode": "021000021",
  "inputType": "KEYED"
}
```

## 🛠 Example Response for Success

```json
{
  "id": "123456",
  "name": "John Doe",
  "accountNumber": "****7890",
  "accountType": "PERSONAL_CHECKING",
  "routingNumber": "****0021",
  "phone": "9876543210",
  "country": "US",
  "default": true
}
```

## Error Handling

| Error Code | Message | Solution | 
| --- | --- | --- |
| `PMT-400` | `"routingNumber is invalid."` | Ensure routing number is exactly of 9 digits
| `401` | `"Unauthorized"` | Check the access token |
| `400` | `"Invalid Request"` | Validate the parameters


## Example Usage in Go

As show in the file [`create_bank_account.go`](/examples/create_bank_account.go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func main() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Your Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	accountData := models.BankAccount{
		Name:          "Tony Stark",
		AccountNumber: "123456789343274",
		AccountType:   "PERSONAL_SAVINGS",
		RoutingNumber: "021000021",
		Phone:         "1234567890",
		BankCode:      "021000021",
	}

	customerID := "12345" // Get from API

	createdAccount, err := apiClient.CreateBankAccount(customerID, accountData)
	if err != nil {
		log.Fatalf("Failed to create bank account: %v", err)
	}

	fmt.Printf("Created Bank Account: %+v\n", createdAccount)
}
```


## References

- [BankAccount APIs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/bankaccounts)