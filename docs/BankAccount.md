# Bank Account API - QuickBooks Intuit SDK

This document provides details on how to create a **Bank Account** using the Intuit API in our SDK.

---

## 🌐 API Endpoint

### **1. Create a Bank Account**

**POST** `/quickbooks/v4/customers/{customer-id}/bank-accounts`

This Go-SDK has `CreateBankAccount(customerID string, accountData models.BankAccount)` function in the client package to hit that request.

```go
customerID := '< of the customer>'

accoundData := '<account details or request body>`
```

### 📥 BankAccount Model

| Parameter       | Required | Type    | Example Value         | Description                         |
| --------------- | -------- | ------- | --------------------- | ----------------------------------- |
| `name`          | ✅       | String  | `"Tony Stark"`        | Name of account holder              |
| `accountNumber` | ✅       | String  | `"123456789012"`      | 4-17 digit account number           |
| `accountType`   | ✅       | Enum    | `"PERSONAL_CHECKING"` | Type of account (Checking/Savings)  |
| `routingNumber` | ✅       | String  | `"021000021"`         | 9-digit routing number              |
| `phone`         | ✅       | String  | `"9876543210"`        | 10-digit phone number               |
| `default`       | ❌       | Boolean | `true`                | Whether this is the default account |
| `country`       | ❌       | String  | `"US"`                | ISO country code                    |
| `bankCode`      | ❌       | String  | `"021000021"`         | Bank code (varies by country)       |
| `inputType`     | ❌       | Enum    | `"KEYED"`             | Method of input                     |

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

accountData := models.BankAccount{
	Name:          "Thor",
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
    "accountNumber": "xxxxxx3274",
    "accountNumberSHA512": "RC1QvzKG50j1Mti3gWO+DacP21BhJDSo0Pza+l+rD9gUpzDvC2Eba/b6bDbeU7xhpMUuGPnBHyVj/nDoHtM9RA==",
    "accountType": "PERSONAL_SAVINGS",
    "bankCode": "021000021",
    "bankName": "JPMORGAN CHASE BANK, NA",
    "country": "US",
    "created": "2025-03-29T07:23:28Z",
    "default": false,
    "entityId": "12345",
    "entityType": "customers",
    "entityVersion": "0",
    "id": "200102476935504778163274",
    "inputType": "KEYED",
    "name": "Thor",
    "phone": "1234567890",
    "routingNumber": "021000021",
    "updated": "2025-03-29T07:23:28Z",
    "verificationStatus": "NOT_VERIFIED"
}
```

---

### **2. Get Details of a Bank Account**

**GET** `/quickbooks/v4/customers/{customer-id}/bank-accounts/{bankaccount-id}`


### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

customerID := "12345"                        // Get from API
bankAccountID := "200102476935504778163274" // Get from API

bankAccount, err := apiClient.GetBankAccount(customerID, bankAccountID)
```

### 🛠 Example Response for Success

```json
{
    "accountNumber": "xxxxxx3274",
    "accountType": "PERSONAL_SAVINGS",
    "bankCode": "021000021",
    "bankName": "JPMORGAN CHASE BANK, NA",
    "country": "US",
    "created": "2025-03-29T07:23:28Z",
    "default": false,
    "entityId": "12345",
    "entityType": "customers",
    "entityVersion": "0",
    "id": "200102476935504778163274",
    "inputType": "KEYED",
    "name": "Thor",
    "phone": "1234567890",
    "routingNumber": "021000021",
    "updated": "2025-03-29T07:23:28Z",
    "verificationStatus": "NOT_VERIFIED"
}
```

---

### **3. Get List of a Bank Accounts**

**GET** `/quickbooks/v4/customers/<id>/bank-accounts`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

customerID := "12345" // Get from API

bankAccounts, err := apiClient.GetBankAccounts(customerID)
```

`bankAccounts` is an array of accounts with specific customerID

### 🛠 Example Response for Success

```json
{"name":"Thor","accountNumber":"xxxxxx3274","accountType":"PERSONAL_SAVINGS","routingNumber":"021000021","phone":"1234567890","bankCode":"021000021","country":"US","inputType":"KEYED"}

{"name":"Tony Stark","accountNumber":"xxxxxx3238","accountType":"PERSONAL_SAVINGS","routingNumber":"021000021","phone":"7534261890","bankCode":"021000021","country":"US","inputType":"KEYED"}
```

---

### **4. DELETE a Bank Accounts**

**DELETE** `/quickbooks/v4/customers/<id>/bank-accounts`

### 🛠 Example Usage

```go
	apiClient := client.NewClient(accessToken)

	customerID := "12345" // Get from API
	accountID := "67890"  // Replace with actual account ID

	resp, err = apiClient.DeleteBankAccount(customerID, accountID) // Replace with actual account ID
```

### 🛠 Example Response for Success

```json
{
  "message": "Bank account deleted successfully",
	"status":  "200",
}
```

---

## References

-   [BankAccount APIs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/bankaccounts)
