package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func CaptureCharge() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	chargeID := "EWKF2SVGW8S3" // Replace with your charge ID

	capture := models.Capture{
		Amount:      "10.00",
		Description: "Refund for order #12345",
		Context: &models.Context{
			Mobile:      "false",
			IsEcommerce: "true",
			Tax:         "0.00",
		},
	}

	resp, err := apiClient.CaptureCharge(chargeID, capture)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Capture response:", resp)
}
