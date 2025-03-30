package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetDetailsRefund() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	chargeID := "EZXGXZQXGXCF" // Replace with your charge ID
	refundID := "EUATXWTYC2ZL" // Replace with your refund ID

	refund, err := apiClient.GetRefund(chargeID, refundID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Refund details:", refund)
}
