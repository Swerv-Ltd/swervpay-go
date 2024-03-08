package main

import (
	"context"
	"fmt"
	"github.com/swerv-ltd/swervpay-go"
	"os"
)

func testCustomer() {

	ctx := context.Background()

	businessId := os.Getenv("SWERVPAY_BUSINESS_ID")
	secretKey := os.Getenv("SWERVPAY_SECRET_KEY")
	baseUrl := os.Getenv("SWERVPAY_BASE_URL")

	client := swervpay.NewSwervpayClient(&swervpay.SwervpayClientOption{
		BusinessID: businessId,
		SecretKey:  secretKey,
		BaseURL:    baseUrl,
	})

	customers, err := client.Customer.Gets(ctx, &swervpay.PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, customer := range *customers {
		fmt.Printf("%v\n", customer)
	}

	customer, err := client.Customer.Get(ctx, "cus_123456")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", customer)

	newCustomer, err := client.Customer.Create(ctx, &swervpay.CreateCustomerBody{
		Firstname:  "John",
		Lastname:   "Doe",
		Middlename: "Doe",
		Email:      "johndoe@gmail.com",
		Country:    "Nigeria",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Created Customer id: " + newCustomer.ID)
	fmt.Printf("%v\n", newCustomer)

}
