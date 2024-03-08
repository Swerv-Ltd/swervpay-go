package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testFx() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	exchange, err := client.Fx.Exchange(ctx, swervpay.FxBody{
		From:   "USD",
		To:     "NGN",
		Amount: 100,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", exchange)

	rate, err := client.Fx.Rate(ctx, swervpay.FxBody{
		From:   "USD",
		To:     "NGN",
		Amount: 100,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", rate)

}
