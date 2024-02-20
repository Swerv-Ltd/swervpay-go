package swervpay

import (
	"context"
	"net/http"
)

type Transaction struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Amount        int64  `json:"amount"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	Category      string `json:"category"`
	Charges       int64  `json:"charges"`
	CreatedAt     string `json:"created_at"`
	Detail        string `json:"detail"`
	FiatRate      int64  `json:"fiat_rate"`
	ID            string `json:"id"`
	Reference     string `json:"reference"`
	Report        bool   `json:"report"`
	ReportMessage string `json:"report_message"`
	SessionID     string `json:"session_id"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	UpdatedAt     string `json:"updated_at"`
}

type TransactionInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Transaction, error)
	Get(ctx context.Context, id string) (*Transaction, error)
}

type TransactionIntImpl struct {
	client *SwervpayClient
}

var _ TransactionInt = &TransactionIntImpl{}

func (t TransactionIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Transaction, error) {
	path := GenerateURLPath("transactions", query)

	// Prepare request
	req, err := t.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Transaction)

	_, err = t.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (t TransactionIntImpl) Get(ctx context.Context, id string) (*Transaction, error) {
	req, err := t.client.NewRequest(ctx, http.MethodGet, "transactions/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Transaction)

	_, err = t.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
