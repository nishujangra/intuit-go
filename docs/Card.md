# Card API - Intuit Payment SDK

This module handles **card operations** in the Intuit Payment SDK, allowing you to:

✅ **Create a new card** for a customer.  
✅ **Use the card for transactions.**  

## **Create a Card**
### **Endpoint:**
```
POST /quickbooks/v4/customers/{id}/cards
```

### **Request Headers:**
| Header      | Value               | Required |
|------------|--------------------|----------|
| Authorization | `Bearer {access_token}` | ✅ |
| Content-Type | `application/json` | ✅ |
| request-Id  | `Unique UUID per request` | ✅ |

### **Request Body:**
### Card Object Request Body

| Attribute          | Type     | Required | Description |
|--------------------|---------|----------|-------------|
| `number`          | String  | ✅ Yes   | The card number. Only the last four digits are displayed for security. |
| `expMonth`        | Integer | ✅ Yes   | The expiration month of the card (1-12). |
| `expYear`         | Integer | ✅ Yes   | The expiration year of the card (e.g., 2027). |
| `cvc`            | String  | ✅ Yes   | The 3 or 4-digit security code on the back of the card. |
| `name`           | String  | ❌ No    | The name of the cardholder. |
| `default`        | Boolean | ❌ No    | If `true`, this card will be used as the default payment method. |
| `commercialCardCode` | String  | ❌ No    | A code identifying a commercial card. |
| `address`        | String  | ❌ No    | The billing address associated with the card. |
| `isBusiness`     | Boolean | ❌ No    | Indicates if the card is a business card. |




### **Example Request:**
```json
{
  "expMonth": "12", 
  "address": {
    "postalCode": "44112", 
    "city": "Richmond", 
    "streetAddress": "1245 Hana Rd", 
    "region": "VA", 
    "country": "US"
  }, 
  "number": "4408041234567893", 
  "name": "Captain Jack Sparrow", 
  "expYear": "2026"
}
```

### **Example Response:**
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