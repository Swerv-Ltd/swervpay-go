package main

import (
	"context"
	"net/http"
)

type CreatePayoutBody struct {
	Reference     string `json:"reference"`
	AccountNumber string `json:"account_number"`
	Narration     string `json:"narration"`
	BankCode      string `json:"bank_code"`
	Currency      string `json:"currency"`
	Amount        int64  `json:"amount"`
}

type CreatePayoutResponse struct {
	Reference string `json:"reference"`
	Message   string `json:"message"`
}

type PayoutInt interface {
	Get(ctx context.Context, id string) (*Transaction, error)
	Create(ctx context.Context, body *CreatePayoutBody) (*CreatePayoutResponse, error)
}

type PayoutIntImpl struct {
	client *SwervpayClient
}

var _ PayoutInt = &PayoutIntImpl{}

func (p PayoutIntImpl) Get(ctx context.Context, id string) (*Transaction, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, "payouts/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Transaction)

	_, err = p.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (p PayoutIntImpl) Create(ctx context.Context, body *CreatePayoutBody) (*CreatePayoutResponse, error) {
	req, err := p.client.NewRequest(ctx, http.MethodPost, "payouts", body)
	if err != nil {
		return nil, err
	}

	response := new(CreatePayoutResponse)

	_, err = p.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
