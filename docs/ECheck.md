# E-Checks API - QuickBooks Intuit SDK

This document provides details on how to create a **E-Checks** using the Intuit API in our SDK and How to use **E-Checks**

-   ✅ **Create a E-Check**
-   ✅ **Refund a E-Check**
-   ✅ **Get Details About E-Check and its refund**

---

## 🌐 API Endpoint

### **1. Create a E-Check**

**POST** `/quickbooks/v4/payments/echecks`

### E-Check Model

| Field         | Type           | Required | Description                           |
| ------------- | -------------- | -------- | ------------------------------------- |
| `Amount`      | `string`       | ✅ Yes   | The amount to be charged.             |
| `PaymentMode` | `string`       | ✅ Yes   | The payment mode (must be `WEB`).     |
| `BankAccount` | `*BankAccount` | ✅ Yes   | The associated bank account details.  |
| `Description` | `string`       | ❌ No    | A description for the transaction.    |
| `CheckNumber` | `string`       | ❌ No    | The check number for the transaction. |
| `Context`     | `*Context`     | ❌ No    | Additional context information.       |

### Bank Account Model

To Check Bank Account Model, [Click here](BankAccount.md)

### Context Object

| Attribute     | Type   | Required | Description                                 |
| ------------- | ------ | -------- | ------------------------------------------- |
| `mobile`      | string | ❌ No    | Indicates if the charge is made via mobile. |
| `isEcommerce` | string | ❌ No    | Indicates if the transaction is e-commerce. |

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

eCheck := models.ECheck{
	Amount:       "100.00",
	PayementMode: "WEB",
	BankAccount: &models.BankAccount{
		AccountNumber: "123456789343274",
		AccountType:   "PERSONAL_SAVINGS",
		Name:          "Thor",
		RoutingNumber: "021000021",
		Phone:         "1234567890",
		BankCode:      "021000021",
	},
	Description: "Payment for services",
	CheckNumber: "123456",
	Context: &models.Context{
		Mobile:      "true",
		IsEcommerce: "true",
		Tax:         "0.00",
	},
}

check, err := apiClient.CreateECheck(eCheck)
```

### 🛠 Example Response for Success

```json
{
    "amount": "100.00",
    "authCode": "776-697",
    "bankAccount": {
        "accountNumber": "xxxxxxxxxxx3274",
        "accountType": "PERSONAL_SAVINGS",
        "inputType": "KEYED",
        "name": "Thor",
        "phone": "1234567890",
        "routingNumber": "xxxxx0021"
    },
    "checkNumber": "123456",
    "context": {
        "deviceInfo": {}
    },
    "created": "2025-03-31T06:56:31Z",
    "description": "Payment for services",
    "id": "a9xts573",
    "paymentMode": "WEB",
    "status": "PENDING"
}
```

---

### **1. Create a E-Check**

**POST** `/quickbooks/v4/payments/echecks`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)
eCheckID := "a9xts573" // Replace with the actual eCheck ID you want to retrieve

check, err := apiClient.GetECheck(eCheckID)
```

### 🛠 Example Response for Success

```json
{
    "amount": "100.00",
    "authCode": "776-697",
    "bankAccount": {
        "accountNumber": "xxxxxxxxxxx3274",
        "accountType": "PERSONAL_SAVINGS",
        "inputType": "KEYED",
        "name": "Thor Thor",
        "phone": "1234567890",
        "routingNumber": "xxxxx0021"
    },
    "checkNumber": "123456",
    "context": {
        "deviceInfo": {}
    },
    "created": "2025-03-31T06:56:31Z",
    "description": "Payment for services",
    "id": "a9xts573",
    "paymentMode": "WEB",
    "status": "PENDING"
}
```

---

## References

-   [E-Check APIs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/echecks)
