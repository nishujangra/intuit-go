package main

import (
	"fmt"
	"log"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func CreateBankAccount() {
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
