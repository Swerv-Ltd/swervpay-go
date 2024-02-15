package main

func main() {
	clientOptions := &SwervpayClientOption{
		BusinessID: "business123",
		SecretKey:  "secretKey123",
		Sandbox:    true,
		Timeout:    30000,
		Version:    "v1",
		BaseURL:    "https://api.swervpay.co/v1",
	}

	swervpayClient := NewSwervpayClient(clientOptions)

}
