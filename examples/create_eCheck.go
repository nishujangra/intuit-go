package main

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func CreateECheck() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Your Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	eCheck := models.ECheck{
		Amount:       "100.00",
		PayementMode: "WEB",
		BankAccount: &models.BankAccount{
			AccountNumber: "123456789343274",
			AccountType:   "PERSONAL_SAVINGS",
			Name:          "Thor",
			RoutingNumber: "021000021",
			Phone:         "1234567890",
			BankCode:      "021000021",
		},
		Description: "Payment for services",
		CheckNumber: "123456",
		Context: &models.Context{
			Mobile:      "true",
			IsEcommerce: "true",
			Tax:         "0.00",
		},
	}

	check, err := apiClient.CreateECheck(eCheck)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(check, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(jsonData))
}
