package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetCharge() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	charge := models.Charge{
		Currency: "USD",
		Amount:   "100.00",
		Context: models.Context{
			Mobile:      "false",
			IsEcommerce: "false",
		},
		Capture:     "true",
		Description: "Test charge",
		CardOnFile:  false,
		Card: models.Card{
			Number:   "4111111111111111",
			CardType: "VISA",
			ExpMonth: "12",
			ExpYear:  "2025",
			CVC:      "123",
			Name:     "Tony Stark",
		},
	}

	createdCharge, err := apiClient.CreateCharge(charge)
	if err != nil {
		fmt.Println("Error creating charge:", err)
		return
	}

	fmt.Println("Charge created successfully:", createdCharge)
}
