package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetDetailsBankAccount() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Your Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	customerID := "12345"                       // Get from API
	bankAccountID := "200102476935504778163274" // Get from API

	bankAccount, err := apiClient.GetBankAccount(customerID, bankAccountID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bank account details:", bankAccount)
}
