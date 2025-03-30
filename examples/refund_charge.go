package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func RefundCharge() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	chargeID := "EZXGXZQXGXCF" // Replace with your charge ID

	refund := models.Refund{
		Amount:      "10.00",
		Description: "Refund for order #12345",
		Context: &models.Context{
			Mobile:      "false",
			IsEcommerce: "true",
			Tax:         "0.00",
		},
	}

	resp, err := apiClient.RefundCharge(chargeID, refund)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Refund response:", resp)
}
