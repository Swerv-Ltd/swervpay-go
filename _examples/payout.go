package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testPayout() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	newPayout, err := client.Payout.Create(ctx, &swervpay.CreatePayoutBody{
		BankCode:      "044",
		AccountNumber: "0690000031",
		Currency:      "NGN",
		Reference:     "123456",
		Amount:        1000,
		Narration:     "Test Payout",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", newPayout)

	payout, err := client.Payout.Get(ctx, newPayout.Reference)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", payout)

}
