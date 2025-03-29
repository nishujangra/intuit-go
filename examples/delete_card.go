package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func DeleteCard() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	customerID := "12345"              // Get from API
	cardId := "xxxxxxxxxxxxxxx7361111" // Get from API

	_, err = apiClient.DeleteCard(customerID, cardId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Card deleted successfully")
}
