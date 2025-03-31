package main

import (
	"encoding/json"
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func GetECheck() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("<Your Auth Code>")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	eCheckID := "a9xts573" // Replace with the actual eCheck ID you want to retrieve

	check, err := apiClient.GetECheck(eCheckID)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(check, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(jsonData))
}
