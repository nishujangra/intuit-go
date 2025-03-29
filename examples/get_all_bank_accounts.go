package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetAllBankAccounts() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	customerID := "12345" // Get from API

	bankAccounts, err := apiClient.GetBankAccounts(customerID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, account := range bankAccounts {
		accountData, err := json.Marshal(account)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(accountData))
	}
}
