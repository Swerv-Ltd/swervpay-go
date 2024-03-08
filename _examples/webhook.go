package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testWebhook() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	wbRes, err := client.Webhook.Test(ctx, "wh_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", wbRes)

	logRes, err := client.Webhook.Retry(ctx, "tri_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", logRes)

}
