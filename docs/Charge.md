# Charge API - Intuit Payment SDK

This module handles **card operations** in the Intuit Payment SDK, allowing you to:

✅ **Create a new Charge** for a customer.  
✅ **Use the Charge for transactions.**

---

## 🌐 API Endpoint

### **1. Create a Charge**

**POST** `/quickbooks/v4/payments/charges`

### Charge Model

| Attribute     | Type   | Required | Description                                                     |
| ------------- | ------ | -------- | --------------------------------------------------------------- |
| `token`       | string | ✅ Yes   | Unique token for the charge.                                    |
| `currency`    | string | ✅ Yes   | Currency code (e.g., "USD").                                    |
| `amount`      | string | ✅ Yes   | Transaction amount.                                             |
| `context`     | object | ✅ Yes   | Context of the transaction (see below).                         |
| `capture`     | bool   | ❌ No    | Whether to capture the charge immediately. Defaults to `false`. |
| `description` | string | ❌ No    | Additional description for the charge.                          |
| `cardOnFile`  | bool   | ❌ No    | Whether the card is stored for future use.                      |
| `card`        | object | ❌ No    | Card details (see below).                                       |

### Context Object

| Attribute     | Type   | Required | Description                                 |
| ------------- | ------ | -------- | ------------------------------------------- |
| `mobile`      | string | ❌ No    | Indicates if the charge is made via mobile. |
| `isEcommerce` | string | ❌ No    | Indicates if the transaction is e-commerce. |

### Card Object

Refer to [Card.md](Card.md) for this object.

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

charge := models.Charge{
	Currency: "USD",
	Amount:   "100.00",
	Context: models.Context{
		Mobile:      "false",
		IsEcommerce: "false",
	},
	Capture:     true,
	Description: "Test charge",
	CardOnFile:  false,
	Card: models.Card{
		Number:   "4111111111111111",
		CardType: "VISA",
		ExpMonth: "12",
		ExpYear:  "2025",
		CVC:      "123",
		Name:     "Tony Stark",
	},
}

chargeCreated, err := apiClient.CreateCharge(charge)
```

### 🛠 Example Response for Success

```json
{
    "amount": "100.00",
    "authCode": "696699",
    "avsStreet": "Pass",
    "avsZip": "Pass",
    "capture": true,
    "card": {
        "cardType": "Visa",
        "cvc": "xxx",
        "expMonth": "12",
        "expYear": "2025",
        "name": "Tony Stark",
        "number": "xxxxxxxxxxxx1111"
},
    "cardSecurityCodeMatch": "NotAvailable",
    "context": {
        "clientTransID": "a9xts389",
        "deviceInfo": {},
        "isEcommerce": false,
        "merchantAccountNumber": "9999996176237953",
        "mobile": false,
        "paymentGroupingCode": "5",
        "paymentStatus": "Completed",
        "reconBatchID": "420250329 1Q13099999996176237953AUTO04",
        "recurring": false,
        "txnAuthorizationStamp": "1743278965"
    },
    "created": "2025-03-29T20:09:25Z",
    "currency": "USD",
    "description": "Test charge",
    "id": "EZXGXZQXGXCF",
    "status": "CAPTURED"
}
```

---
