# Bank Account API - QuickBooks Intuit SDK

This document provides details on how to create a **Bank Account** using the Intuit API in our SDK.

## 🌐 API Endpoint

### **1. Create a Bank Account**
**POST** `/quickbooks/v4/customers/{customer-id}/bank-accounts`

This Go-SDK has `CreateBankAccount(customerID string, accountData models.BankAccount)` function in the client package to hit that request.

```go
customerID := '< of the customer>'

accoundData := '<account details or request body>`
```

### 📥 BankAccount Model
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


### 🛠 Example Usage
```go
accountData := models.BankAccount{
	Name:          "Tony Stark",
	AccountNumber: "123456789343274",
	AccountType:   "PERSONAL_SAVINGS",
	RoutingNumber: "021000021",
	Phone:         "1234567890",
	BankCode:      "021000021",
}

customerID := "12345" // Get from API

bankAccount, err := apiClient.CreateBankAccount(customerID, accountData)
```

### 🛠 Example Response for Success

```json
{
  "id": "123456",
  "name": "Tony Stark",
  "accountNumber": "**********3274",
  "accountType": "PERSONAL_SAVINGS",
  "routingNumber": "****0021",
  "phone": "1234567890",
  "country": "US",
  "default": true
}
```

---

### **2. Get Details of a Bank Account**
**GET** `/quickbooks/v4/customers/{customer-id}/bank-accounts/{bankaccount-id}`

After making a api request on this endpoint example response will be like 

### 🛠 Example Response for Success
```json
{
  "updated": "2025-03-28T23:25:19Z", 
  "name": "Richard Jones", 
  "accountNumber": "XXXXXXXXXXX3274", 
  "default": false, 
  "created": "2025-03-28T16:05:39Z", 
  "inputType": "KEYED", 
  "phone": "6047296480", 
  "accountType": "PERSONAL_SAVINGS", 
  "routingNumber": "XXXXX0021", 
  "id": "200161921532106731364534"
}
```



## References

- [BankAccount APIs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/bankaccounts)