package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func VoidCharge() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Your Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	chargeRequestID := "EWKF2SVGW8S3"

	voidCharge, err := apiClient.VoidCharge(chargeRequestID)
	if err != nil {
		fmt.Println("Error voiding charge:", err)
		return
	}

	fmt.Println("Void Charge Response:")
	fmt.Println(voidCharge)
}
