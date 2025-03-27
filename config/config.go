package config

import (
	"github.com/spf13/viper"
)

func viperInit() {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	viper.ReadInConfig()
}

func GetClientID() string {
	viperInit()
	return viper.GetString("INTUIT.INTUIT-CREDENTIALS.CLIENT_ID")
}

func GetClientSecret() string {
	viperInit()
	return viper.GetString("INTUIT.INTUIT-CREDENTIALS.CLIENT_SECRET")
}

func GetPaymentsBaseURL() string {
	viperInit()
	return viper.GetString("INTUIT.PAYMENTS.SANDBOX.BASE_URL")
}

func GetPaymentsAuthURL() string {
	viperInit()
	return viper.GetString("INTUIT.PAYMENTS.SANDBOX.AUTH_URL")
}
