package main

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetAllCards() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

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
}
