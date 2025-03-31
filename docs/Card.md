# Card API - Intuit Payment SDK

This module handles **card operations** in the Intuit Payment SDK, allowing you to:

- ✅ **Create a new card** for a customer.  
- ✅ **Get Details of the card** for a customer
- ✅ **Get list of the cards** for a customer
- ✅ **Delete a card** for a customer

---

## 🌐 API Endpoint

### **1. Create a Card**

**POST** `/quickbooks/v4/customers/{id}/cards`

### Card Model

| Attribute            | Type    | Required | Description                                                            |
| -------------------- | ------- | -------- | ---------------------------------------------------------------------- |
| `number`             | String  | ✅ Yes   | The card number. Only the last four digits are displayed for security. |
| `expMonth`           | Integer | ✅ Yes   | The expiration month of the card (1-12).                               |
| `expYear`            | Integer | ✅ Yes   | The expiration year of the card (e.g., 2027).                          |
| `cvc`                | String  | ✅ Yes   | The 3 or 4-digit security code on the back of the card.                |
| `name`               | String  | ❌ No    | The name of the cardholder.                                            |
| `default`            | Boolean | ❌ No    | If `true`, this card will be used as the default payment method.       |
| `commercialCardCode` | String  | ❌ No    | A code identifying a commercial card.                                  |
| `address`            | String  | ❌ No    | The billing address associated with the card.                          |
| `isBusiness`         | Boolean | ❌ No    | Indicates if the card is a business card.                              |

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

cardData := models.Card{
	Number:   "4111111111111111",
	CardType: "VISA",
	ExpMonth: "12",
	ExpYear:  "2025",
	CVC:      "123",
	Name:     "Tony Stark",
}

customerID := "12345" // Get from API

createdCard, err := apiClient.CreateCard(customerID, cardData)
```

### 🛠 Example Response for Success

```json
{
    "updated": "2025-03-28T22:24:33Z",
    "name": "Captain Jack Sparrow",
    "created": "2025-03-28T22:24:33Z",
    "default": false,
    "number": "xxxxxxxxxxxx7893",
    "expMonth": "12",
    "address": {
        "postalCode": "44112",
        "city": "Richmond",
        "region": "VA",
        "streetAddress": "1245 Hana Rd",
        "country": "US"
    },
    "expYear": "2026",
    "isBusiness": false,
    "id": "101101015602106732027893"
}
```

---

### **2. Get Details of a Card**

**GET** `/quickbooks/v4/customers/{customer id}/cards/{Card ID}`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

customerID := "12345"                // Get from API
cardId := "101143461745504777361111" // Get from API

card, err := apiClient.GetCard(customerID, cardId)
if err != nil {
	fmt.Println(err)
  return
}
```

### 🛠 Example Response for Success

```json
{
    "cardType": "VISA",
    "created": "2025-03-28T18:54:54Z",
    "cvcVerification": { "date": "2025-03-28T18:54:54Z", "result": "NOT_DONE" },
    "default": false,
    "entityId": "12345",
    "entityType": "customers",
    "entityVersion": "0",
    "expMonth": "12",
    "expYear": "2025",
    "id": "101143461745504777361111",
    "isBusiness": false,
    "isLevel3Eligible": false,
    "name": "Tony Stark",
    "number": "xxxxxxxxxxxx1111",
    "numberSHA512": "MBDxroON6A5ittif8c/zMyPsbRq75gYiT+FjP0o8rfOwlBQ76sfR9L+56Mjk4qkbG8nwSkGNIb0fb/Z3rfgIlA==",
    "status": "ACTIVE",
    "updated": "2025-03-28T18:54:54Z",
    "zeroDollarVerification": { "status": "NOT_VERIFIED" }
}
```

---

### **3. Get List of Cards**

**GET** `/quickbooks/v4/customers/{customer-id}/cards`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

customerID := "12345" // Get from API

cards, err := apiClient.GetCards(customerID)
if err != nil {
	fmt.Println(err)
	return
}

for _, card := range cards {
	jsonData, err := json.Marshal(card)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
```

### 🛠 Example Response for Success

```json
[
    {
        "cardType": "VISA",
        "created": "2025-03-28T18:54:54Z",
        "cvcVerification": {
            "date": "2025-03-28T18:54:54Z",
            "result": "NOT_DONE"
        },
        "default": false,
        "entityId": "12345",
        "entityType": "customers",
        "entityVersion": "0",
        "expMonth": "12",
        "expYear": "2025",
        "id": "101143461745504777361111",
        "isBusiness": false,
        "isLevel3Eligible": false,
        "name": "Tony Stark",
        "number": "xxxxxxxxxxxx1111",
        "numberSHA512": "MBDxroON6A5ittif8c/zMyPsbRq75gYiT+FjP0o8rfOwlBQ76sfR9L+56Mjk4qkbG8nwSkGNIb0fb/Z3rfgIlA==",
        "status": "ACTIVE",
        "updated": "2025-03-28T18:54:54Z",
        "zeroDollarVerification": { "status": "NOT_VERIFIED" }
    }
]
```

---

### **4. Delete a Card**

**DELETE** `/quickbooks/v4/customers/{customer id}/cards/{Card ID}`

### 🛠 Example Usage

```go
apiClient := client.NewClient(accessToken)

customerID := "12345"              // Get from API
cardId := "xxxxxxxxxxxxxxx7361111" // Get from API

_, err = apiClient.DeleteCard(customerID, cardId)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println("Card deleted successfully")

```

### 🛠 Example Response for Success

```json
{
    "status": "success",
    "message": "Card deleted successfully",
    "code": "200"
}
```

---

## References

-   [Card APIs](https://developer.intuit.com/app/developer/qbpayments/docs/api/resources/all-entities/cards)
