package swervpay

import (
	"context"
	"net/http"
)

// Bank represents a bank in the Swervpay system.
type Bank struct {
	Code string `json:"bank_code"` // Code is the unique identifier for the bank.
	Name string `json:"bank_name"` // Name is the name of the bank.
}

// ResolveAccountNumber represents the response from the Swervpay API when resolving an account number.
type ResolveAccountNumber struct {
	AccountNumber string `json:"account_number"` // AccountNumber is the account number that was resolved.
	BankCode      string `json:"bank_code"`      // BankCode is the code of the bank the account belongs to.
	BankName      string `json:"bank_name"`      // BankName is the name of the bank the account belongs to.
	AccountName   string `json:"account_name"`   // AccountName is the name of the account holder.
}

// ResolveAccountNumberBody represents the request body when resolving an account number.
type ResolveAccountNumberBody struct {
	AccountNumber string `json:"account_number"` // AccountNumber is the account number to resolve.
	BankCode      string `json:"bank_code"`      // BankCode is the code of the bank the account belongs to.
}

// OtherInt is an interface for interacting with the Swervpay API.
type OtherInt interface {
	// Banks retrieves a list of all banks in the Swervpay system.
	Banks(ctx context.Context) ([]*Bank, error)
	// ResolveAccountNumber resolves an account number in the Swervpay system.
	ResolveAccountNumber(ctx context.Context, body ResolveAccountNumberBody) (*ResolveAccountNumber, error)
}

// OtherIntImpl is an implementation of the OtherInt interface.
type OtherIntImpl struct {
	client *SwervpayClient // client is the Swervpay client used to interact with the API.
}

// Verify that OtherIntImpl implements the OtherInt interface.
var _ OtherInt = &OtherIntImpl{}

// Banks retrieves a list of all banks in the Swervpay system.
// https://docs.swervpay.co/api-reference/others/get-banks
func (o OtherIntImpl) Banks(ctx context.Context) ([]*Bank, error) {

	req, err := o.client.NewRequest(ctx, http.MethodGet, "banks", nil)
	if err != nil {
		return nil, err
	}

	response := []*Bank{}

	_, err = o.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// ResolveAccountNumber resolves an account number in the Swervpay system.
// https://docs.swervpay.co/api-reference/others/resolve-account-number
func (o OtherIntImpl) ResolveAccountNumber(ctx context.Context, body ResolveAccountNumberBody) (*ResolveAccountNumber, error) {

	req, err := o.client.NewRequest(ctx, http.MethodPost, "resolve-account-number", body)
	if err != nil {
		return nil, err
	}

	response := new(ResolveAccountNumber)

	_, err = o.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
