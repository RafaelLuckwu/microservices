package config

import "os"

func GetApplicationPort() string {
	p := os.Getenv("APPLICATION_PORT")
	if p == "" { return ":3000" }
	return ":" + p
}

func GetDataSourceURL() string { return os.Getenv("DATA_SOURCE_URL") }

func GetPaymentServiceUrl() string {
	url := os.Getenv("PAYMENT_SERVICE_URL")
	if url == "" { return "localhost:3001" }
	return url
}
