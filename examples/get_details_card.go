package main

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetDetailsCard() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	customerID := "12345"                // Get from API
	cardId := "xxxxxxxxxxxxx04777361111" // Get from API

	card, err := apiClient.GetCard(customerID, cardId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Card Details:")
	jsonData, err := json.Marshal(card)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
