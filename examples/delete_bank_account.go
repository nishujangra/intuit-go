package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func DeleteBankAccount() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	customerID := "12345" // Get from API
	accountID := "67890"  // Replace with actual account ID

	_, err = apiClient.DeleteBankAccount(customerID, accountID) // Replace with actual account ID
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bank account deleted successfully")
}
