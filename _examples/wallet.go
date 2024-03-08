package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testWallet() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	wallets, err := client.Wallet.Gets(ctx, &swervpay.PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, wallet := range *wallets {
		fmt.Printf("%v\n", wallet)
	}

	//wallet, err := client.Wallet.Get(ctx, walletId)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%v\n", wallet)

}
