package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testOther() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	banks, err := client.Other.Banks(ctx)

	if err != nil {
		panic(err)
	}

	for _, bank := range *banks {
		fmt.Printf("%v\n", bank)
	}

	resolveAccount, err := client.Other.ResolveAccountNumber(ctx, swervpay.ResolveAccountNumberBody{
		BankCode:      "044",
		AccountNumber: "0690000031",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", resolveAccount)

}
