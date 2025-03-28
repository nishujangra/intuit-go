package main

import (
	"fmt"

	"github.com/nishujangra/intuit-go/config"
	"github.com/nishujangra/intuit-go/models"
	"github.com/nishujangra/intuit-go/pkg/auth"
	"github.com/nishujangra/intuit-go/pkg/client"
)

func CreateCard() {
	authClient := auth.NewAuthClient(config.GetClientID(), config.GetClientSecret(), config.GetPaymentsBaseURL())

	accessToken, err := authClient.GetToken("AB11743188166zAoePRJcZpSkNcAgQvByhtkz5uz57LTC2F81c")
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.NewClient(accessToken)

	cardData := models.Card{
		Number:   "4111111111111111",
		CardType: "VISA",
		ExpMonth: "12",
		ExpYear:  "2025",
		CVC:      "123",
		Name:     "Tony Stark",
	}

	customerID := "12345" // Get from API

	createdCard, err := apiClient.CreateCard(customerID, cardData)
	if err != nil {
		fmt.Printf("Failed to create card: %v\n", err)
		return
	}
	fmt.Printf("Created Card: %+v\n", createdCard)
}
