package swervpay

import (
	"context"
	"net/http"
)

type Wallet struct {
	AccountName    string `json:"account_name"`
	AccountNumber  string `json:"account_number"`
	AccountType    string `json:"account_type"`
	Balance        int64  `json:"balance"`
	BankAddress    string `json:"bank_address"`
	BankCode       string `json:"bank_code"`
	BankName       string `json:"bank_name"`
	CreatedAt      string `json:"created_at"`
	ID             string `json:"id"`
	IsBlocked      bool   `json:"is_blocked"`
	Label          string `json:"label"`
	PendingBalance int64  `json:"pending_balance"`
	Reference      string `json:"reference"`
	RoutingNumber  string `json:"routing_number"`
	TotalReceived  int64  `json:"total_received"`
	UpdatedAt      string `json:"updated_at"`
}

type WalletInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Wallet, error)
	Get(ctx context.Context, id string) (*Wallet, error)
}

type WalletIntImpl struct {
	client *SwervpayClient
}

var _ WalletInt = &WalletIntImpl{}

func (w WalletIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Wallet, error) {

	path := GenerateURLPath("wallets", query)

	// Prepare request
	req, err := w.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Wallet)

	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (w WalletIntImpl) Get(ctx context.Context, id string) (*Wallet, error) {
	// Prepare request
	req, err := w.client.NewRequest(ctx, http.MethodGet, "wallets/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
