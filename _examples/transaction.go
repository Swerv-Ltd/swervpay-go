package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testTransaction() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	transactions, err := client.Transaction.Gets(ctx, &swervpay.PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, transaction := range *transactions {
		fmt.Printf("%v\n", transaction)
	}

	transaction, err := client.Transaction.Get(ctx, "txn_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", transaction)
}
