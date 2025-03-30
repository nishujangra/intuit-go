# 🏦 Intuit Payment SDK (Go)

An open-source **Golang SDK** for integrating with **Intuit's QuickBooks Payments API**, allowing:

- **Authentication using OAuth2**
- **Bank Account Creation**
- **Card Transactions**

## 🚀 Features
✔ **Bank Account Management**: Create and manage bank accounts via QuickBooks API.  
✔ **Secure Authentication**: Uses OAuth2 bearer tokens for API access.  
✔ **Error Handling**: Includes structured error handling for API requests.


## Installation
```sh
go get github.com/nishujangra/intuit-go
```

---

## 🔑 Setup

### 1. Obtain API Crendtials from [Intuit Developer Portal](https://developer.intuit.com/)

### 2. Setup the configuration file

Create a `config.json` file in the main directory of the project you are using this sdk in.

```sh
touch config.json
```

Fill the `config.json` file with

```json
{
    "INTUIT": {
        "INTUIT-CREDENTIALS": {
            "CLIENT_ID": "<your client id>",
            "CLIENT_SECRET": "<your client secret>"
        },
        "PAYMENTS": {
            "PRODUCTION": {
                "BASE_URL": "https://api.intuit.com/",
                "AUTH_URL": "https://oauth.platform.intuit.com/"
            },
            "SANDBOX": {
                "BASE_URL": "https://sandbox.api.intuit.com/",
                "AUTH_URL": "https://oauth.platform.intuit.com/"
            }
        },
    }
}
```

### 3. Make API client

```go
authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

accessToken, err := authClient.GetToken("<Auth Token>")
if err != nil {
	fmt.Println(err)
}

apiClient := client.NewClient(accessToken)
```

---

## Usage
Check the [examples](examples/) for how to use this SDK.

## Related Links

- [How to setup authentication](docs/Auth.md)
- [Bank Account](docs/BankAccount.md)
- [Crate Card](docs/Card.md)
- [Charges, Refund](docs/Charge.md)

---

## 🛠 Contributing

👨‍💻 Contributions are welcome! Follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit changes (`git commit -m "Added new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a Pull Request

## LICENSE

This project comes under the [License](LICENSE.md)

## 💡 Support
For any issues or feature requests, open an issue in [GitHub Issues](https://github.com/nishujangra/intuit-go/issues). 🚀