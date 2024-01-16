// Package swyftpay provides functionality for interacting with Swyftpay API.
package swyftpay

// SwyftpayClientOption represents the options for SwyftpayClient.
type SwyftpayClientOption struct {
	BusinessID string
	SecretKey  string
	Sandbox    bool
	Timeout    int
	Version    string
	BaseURL    string
}

// SwyftpayClient represents a client for interacting with Swyftpay API.
type SwyftpayClient struct {
	options *SwyftpayClientOption
}

// NewSwyftpayClient creates a new SwyftpayClient with the given options.
func NewSwyftpayClient(options *SwyftpayClientOption) *SwyftpayClient {
	return &SwyftpayClient{
		options: options,
	}
}

// SwyftpayApiClient represents a client for Swyftpay API.
type SwyftpayApiClient struct{}

// NewSwyftpayApiClient creates a new SwyftpayApiClient.
func NewSwyftpayApiClient() *SwyftpayApiClient {
	return &SwyftpayApiClient{}
}
