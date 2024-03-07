package swervpay

import (
	"context"
	"net/http"
)

// Transaction represents a transaction with all its details.
type Transaction struct {
	AccountName   string `json:"account_name"`   // The name of the account
	AccountNumber string `json:"account_number"` // The number of the account
	Amount        int64  `json:"amount"`         // The amount of the transaction
	BankCode      string `json:"bank_code"`      // The code of the bank
	BankName      string `json:"bank_name"`      // The name of the bank
	Category      string `json:"category"`       // The category of the transaction
	Charges       int64  `json:"charges"`        // The charges of the transaction
	CreatedAt     string `json:"created_at"`     // The creation date of the transaction
	Detail        string `json:"detail"`         // The detail of the transaction
	FiatRate      int64  `json:"fiat_rate"`      // The fiat rate of the transaction
	ID            string `json:"id"`             // The ID of the transaction
	Reference     string `json:"reference"`      // The reference of the transaction
	Report        bool   `json:"report"`         // The report status of the transaction
	ReportMessage string `json:"report_message"` // The report message of the transaction
	SessionID     string `json:"session_id"`     // The session ID of the transaction
	Status        string `json:"status"`         // The status of the transaction
	Type          string `json:"type"`           // The type of the transaction
	UpdatedAt     string `json:"updated_at"`     // The update date of the transaction
}

// TransactionInt is an interface that defines the methods for transactions.
type TransactionInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Transaction, error) // Gets a list of transactions
	Get(ctx context.Context, id string) (*Transaction, error)                   // Gets a single transaction
}

// TransactionIntImpl is the implementation of the TransactionInt interface.
type TransactionIntImpl struct {
	client *SwervpayClient // The client used to make requests
}

// Verify that TransactionIntImpl implements TransactionInt.
var _ TransactionInt = &TransactionIntImpl{}

// Gets retrieves a list of transactions.
// https://docs.swervpay.co/api-reference/transactions/get-all-transactions
func (t TransactionIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Transaction, error) {
	path := GenerateURLPath("transactions", query)

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

// Get retrieves a single transaction.
// https://docs.swervpay.co/api-reference/transactions/get
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
