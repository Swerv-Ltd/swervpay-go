package main

import (
	"github.com/swerv-ltd/swervpay-go/swyftpay"
)

func main() {
	clientOptions := &swyftpay.SwyftpayClientOption{
		BusinessID: "business123",
		SecretKey:  "secretKey123",
		Sandbox:    true,
		Timeout:    30000,
		Version:    "v1",
		BaseURL:    "https://api.swervpay.co/v1",
	}

	swyftpayClient := swyftpay.NewSwyftpayClient(clientOptions)
	//swyftpayApiClient := swyftpay.NewSwyftpayApiClient()

	swyftpayClient.PrintOptions()
}
