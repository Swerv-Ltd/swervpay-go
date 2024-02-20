package swervpay

import (
	"context"
	"net/http"
)

type Bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResolveAccountNumber struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
}

type ResolveAccountNumberBody struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

type OtherInt interface {
	Banks(ctx context.Context) (*[]Bank, error)
	ResolveAccountNumber(ctx context.Context, body ResolveAccountNumberBody) (*ResolveAccountNumber, error)
}

type OtherIntImpl struct {
	client *SwervpayClient
}

var _ OtherInt = &OtherIntImpl{}

func (o OtherIntImpl) Banks(ctx context.Context) (*[]Bank, error) {
	// Prepare request
	req, err := o.client.NewRequest(ctx, http.MethodGet, "banks", nil)
	if err != nil {
		return nil, err
	}

	response := new([]Bank)

	_, err = o.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (o OtherIntImpl) ResolveAccountNumber(ctx context.Context, body ResolveAccountNumberBody) (*ResolveAccountNumber, error) {
	// Prepare request
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
