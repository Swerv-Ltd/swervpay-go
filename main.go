package main

import (
	"swervpay-go/swyftpay"
)

func main() {
	clientOptions := &swyftpay.SwyftpayClientOption{
		BusinessID: "business123",
		SecretKey:  "secretKey123",
		Sandbox:    true,
		Timeout:    30000,
		Version:    "v1",
		BaseURL:    "https://swyftpay.com/api",
	}

	swyftpayClient := swyftpay.NewSwyftpayClient(clientOptions)
	//swyftpayApiClient := swyftpay.NewSwyftpayApiClient()

	swyftpayClient.PrintOptions()
}
