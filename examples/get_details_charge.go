package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetDetailsCharge() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	chargeID := "EZXGXZQXGXCF" // Replace with your charge ID

	charge, err := apiClient.GetCharge(chargeID)
	if err != nil {
		fmt.Println("Error getting charge:", err)
		return
	}

	fmt.Println("Charge details:", charge)
}
