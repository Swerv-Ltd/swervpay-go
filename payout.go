package swervpay

import (
	"context"
	"net/http"
)

// CreatePayoutBody represents the request body for creating a payout.
type CreatePayoutBody struct {
	Reference     string  `json:"reference"`      // Unique reference for the payout
	AccountNumber string  `json:"account_number"` // Account number to send the payout to
	Narration     string  `json:"narration"`      // Description of the payout
	BankCode      string  `json:"bank_code"`      // Code of the bank for the account
	Currency      string  `json:"currency"`       // Currency of the payout
	Amount        float64 `json:"amount"`         // Amount of the payout
}

// CreatePayoutResponse represents the response from creating a payout.
type CreatePayoutResponse struct {
	Reference string `json:"reference"` // Unique reference for the payout
	ID        string `json:"id"`        // ID of the payout
	Message   string `json:"message"`   // Message indicating the status of the payout
}

// PayoutInt is an interface for managing payouts.
type PayoutInt interface {
	// Get retrieves a payout by its ID.
	Get(ctx context.Context, id string) (*Transaction, error)
	// Create creates a new payout with the provided body.
	Create(ctx context.Context, body *CreatePayoutBody) (*CreatePayoutResponse, error)
}

// PayoutIntImpl is an implementation of the PayoutInt interface.
type PayoutIntImpl struct {
	client *SwervpayClient // Client used to make requests
}

// Verify that PayoutIntImpl implements PayoutInt.
var _ PayoutInt = &PayoutIntImpl{}

// Get retrieves a payout by its ID.
// https://docs.swervpay.co/api-reference/payouts/get
func (p PayoutIntImpl) Get(ctx context.Context, id string) (*Transaction, error) {
	// Create a new request to get a payout.
	req, err := p.client.NewRequest(ctx, http.MethodGet, "payouts/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Transaction)

	// Perform the request and populate the response.
	_, err = p.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	// Return the retrieved payout.
	return response, nil
}

// Create creates a new payout with the provided body.
// https://docs.swervpay.co/api-reference/payouts/create
func (p PayoutIntImpl) Create(ctx context.Context, body *CreatePayoutBody) (*CreatePayoutResponse, error) {
	// Create a new request to create a payout.
	req, err := p.client.NewRequest(ctx, http.MethodPost, "payouts", body)
	if err != nil {
		return nil, err
	}

	response := new(CreatePayoutResponse)

	// Perform the request and populate the response.
	_, err = p.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	// Return the response from creating the payout.
	return response, nil
}
