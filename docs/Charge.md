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

### **2. Get a Charge**

**GET** `/quickbooks/v4/payments/charges/{charge-id}`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

chargeId := "EZXGXZQHXXCF" // Replace with your charge ID

charge, err := apiClient.GetCharge(chargeId)
```

### 🛠 Example Response for Success

```json
{
    "amount": "100.00",
    "appType": "7716935554825128052",
    "authCode": "696699",
    "avsStreet": "Pass",
    "avsZip": "Pass",
    "capture": true,
    "card": {
        "address": { "country": "USA" },
        "cardType": "Visa",
        "expMonth": "12",
        "expYear": "2025",
        "name": "Tony Stark",
        "number": "xxxxxxxxxxxx1111"
    },
    "cardSecurityCodeMatch": "NotAvailable",
    "context": { "isEcommerce": false, "mobile": false, "recurring": false },
    "created": "2025-03-29T20:09:26Z",
    "currency": "USD",
    "description": "Test charge",
    "id": "EZXGXZQXGXCF",
    "status": "CAPTURED"
}
```

---

### **3. Refund a charge**

**POST** `v4/payments/charges/{charge-id}/refunds`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

chargeId := "EZXGXZQHXXCF" // Replace with your charge ID

refund := models.Refund{
	Amount:      "10.00",
	Description: "Refund for order #12345",
	Context: models.Context{
		Mobile:      "false",
		IsEcommerce: "true",
		Tax:         "0.00",
	},
}

refund, err := apiClient.RefundCharge(chargeID, refund)
```

### 🛠 Example Response for Success

```json
{
    "amount": "10.00",
    "context": {
        "clientTransID": "a9xts3wz",
        "isEcommerce": true,
        "mobile": false,
        "paymentGroupingCode": "5",
        "reconBatchID": "420250330 1Q01449999996176237953AUTO04",
        "recurring": false,
        "tax": "0.00",
        "txnAuthorizationStamp": "1743324244"
    },
    "created": "2025-03-30T08:44:04Z",
    "description": "Refund for order #12345",
    "id": "EFL8WUV0W36R",
    "status": "ISSUED",
    "type": "REFUND"
}
```

---

### **4. Get a refund by ID**

**GET** `/quickbooks/v4/payments/charges/{charge-id}/refunds/{refund-id}`

### 🛠 Example Usage

```go
aapiClient := client.NewClient(accessToken)

chargeID := "EZXGXZQXGXCF" // Replace with your charge ID
refundID := "EUATXWTYC2ZL" // Replace with your refund ID

refund, err := apiClient.GetRefund(chargeID, refundID)
```

### 🛠 Example Response for Success

```json
{
    "amount": "5.00",
    "context": { "isEcommerce": false, "mobile": false, "recurring": false },
    "created": "2025-03-30T07:58:42Z",
    "description": "first refund",
    "id": "EUATXWTYC2ZL",
    "status": "ISSUED",
    "type": "REFUND"
}
```

---

### **5. Capture of Charge**

**POST** `/quickbooks/v4/payments/charges/{charge-id}/capture`

`Note` : To capture a charge, the status of the charge should not be `CAPTURED`. Set `capture:false` while creating charge

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

chargeID := "EWKF2SVGW8S3" // Replace with your charge ID

capture := models.Capture{
	Amount:      "10.00",
	Description: "Refund for order #12345",
	Context: &models.Context{
		Mobile:      "false",
		IsEcommerce: "true",
		Tax:         "0.00",
	},
}

resp, err := apiClient.CaptureCharge(chargeID, capture)
```

### 🛠 Example Response for Success

```json
{
    "amount": "100.00",
    "appType": "7716935554825128052",
    "authCode": "503100",
    "avsStreet": "Pass",
    "avsZip": "Pass",
    "capture": false,
    "captureDetail": {
        "amount": "10.00",
        "context": {
            "isEcommerce": false,
            "mobile": false,
            "recurring": false,
            "tax": "0.00"
        },
        "created": "2025-03-30T08:55:52Z",
        "description": "Refund for order #12345"
    },
    "card": {
        "address": {},
        "cardType": "Visa",
        "expMonth": "12",
        "expYear": "2025",
        "name": "Tony Stark",
        "number": "xxxxxxxxxxxx1111"
    },
    "cardSecurityCodeMatch": "NotAvailable",
    "context": {
        "isEcommerce": false,
        "merchantAccountNumber": "9999996176237953",
        "mobile": false,
        "paymentGroupingCode": "5",
        "reconBatchID": "420250330 1Q01559999996176237953AUTO04",
        "recurring": false,
        "txnAuthorizationStamp": "1743324952"
    },
    "created": "2025-03-30T08:55:26Z",
    "currency": "USD",
    "description": "Test charge",
    "id": "EWKF2SVGW8S3",
    "status": "CAPTURED"
}
```

---

### **6. Void a Charge**

**POST**

`Note:` You would require `request-id` of the charge you want to avoid. I would recommed you to store the `request-id` to your database to reuse them.

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

chargeRequestID := "chrg_3J2Y1Z2eZvKYlo2C7g4v5a0d"

voidCharge, err := apiClient.VoidCharge(chargeRequestID)
```

### 🛠 Example Response for Success

```json
{
    "status": "ISSUED",
    "created": "2025-5-30T10:59:23Z",
    "amount": "1.00",
    "context": {
        "mobile": false,
        "recurring": false
    },
    "type": "VOID",
    "id": "E7UT68A501QI"
}
```

---

## References

-   [Charges API Docs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/charges)
