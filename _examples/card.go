package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testCards() {
	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	// Get all cards
	cards, err := client.Card.Gets(ctx, &swervpay.PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, card := range *cards {
		fmt.Printf("%v\n", card)
	}

	// Get a card
	card, err := client.Card.Get(ctx, "card_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", card)
}
