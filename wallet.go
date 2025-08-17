package swervpay

import (
	"context"
	"net/http"
)

// Wallet represents a user's wallet in the system.
type Wallet struct {
	AccountName    string  `json:"account_name"`    // The name of the account.
	AccountNumber  string  `json:"account_number"`  // The number of the account.
	AccountType    string  `json:"account_type"`    // The type of the account.
	Balance        float64 `json:"balance"`         // The current balance of the wallet.
	BankAddress    string  `json:"bank_address"`    // The address of the bank.
	BankCode       string  `json:"bank_code"`       // The code of the bank.
	BankName       string  `json:"bank_name"`       // The name of the bank.
	CreatedAt      string  `json:"created_at"`      // The creation date of the wallet.
	ID             string  `json:"id"`              // The unique identifier of the wallet.
	IsBlocked      bool    `json:"is_blocked"`      // Indicates if the wallet is blocked.
	Label          string  `json:"label"`           // The label of the wallet.
	PendingBalance float64 `json:"pending_balance"` // The pending balance of the wallet.
	Reference      string  `json:"reference"`       // The reference of the wallet.
	RoutingNumber  string  `json:"routing_number"`  // The routing number of the bank.
	TotalReceived  float64 `json:"total_received"`  // The total amount received in the wallet.
	UpdatedAt      string  `json:"updated_at"`      // The last update date of the wallet.
}

// WalletInt is an interface that defines the methods for managing wallets.
type WalletInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Wallet, error) // Gets a list of wallets.
	Get(ctx context.Context, id string) (*Wallet, error)                   // Gets a specific wallet by its ID.
}

// WalletIntImpl is an implementation of the WalletInt interface.
type WalletIntImpl struct {
	client *SwervpayClient // The client used to make requests.
}

// Verify that WalletIntImpl implements WalletInt.
var _ WalletInt = &WalletIntImpl{}

// Gets retrieves a list of wallets.
// https://docs.swervpay.co/api-reference/wallets/get-all-wallets
func (w WalletIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Wallet, error) {

	path := GenerateURLPath("wallets", query)

	// Prepare the request.
	req, err := w.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*Wallet{}

	// Perform the request and get the response.
	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	// Return the response.
	return response, nil
}

// Get retrieves a specific wallet by its ID.
// https://docs.swervpay.co/api-reference/wallets/get
func (w WalletIntImpl) Get(ctx context.Context, id string) (*Wallet, error) {
	// Prepare the request.
	req, err := w.client.NewRequest(ctx, http.MethodGet, "wallets/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	// Perform the request and get the response.
	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	// Return the response.
	return response, nil
}
