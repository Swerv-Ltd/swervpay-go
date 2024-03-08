package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testCollection() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	// Get all collections
	collections, err := client.Collection.Gets(ctx, &swervpay.PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, collection := range *collections {
		fmt.Printf("%v\n", collection)
	}

	// Get a collection
	collection, err := client.Collection.Get(ctx, "wal_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", collection)

	// Create a collection
	newCollection, err := client.Collection.Create(ctx, &swervpay.CreateCollectionBody{
		Amount:       1000,
		Currency:     "NGN",
		Type:         "ONE_TIME",
		MerchantName: "John Doe",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Created Collection id: " + newCollection.ID)

	fmt.Printf("%v\n", newCollection)

}
