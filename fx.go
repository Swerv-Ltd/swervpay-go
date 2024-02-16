package main

import (
	"context"
	"net/http"
)

type FxBody struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type FxRateResponse struct {
	Rate float64  `json:"rate"`
	From FromOrTo `json:"from"`
	To   FromOrTo `json:"to"`
}

type FromOrTo struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type FxInt interface {
	Rate(ctx context.Context, body FxBody) (*FxRateResponse, error)
	Exchange(ctx context.Context, body FxBody) (*Transaction, error)
}

type FxIntImpl struct {
	client *SwervpayClient
}

var _ FxInt = &FxIntImpl{}

func (f FxIntImpl) Rate(ctx context.Context, body FxBody) (*FxRateResponse, error) {
	req, err := f.client.NewRequest(ctx, http.MethodPost, "fx/rate", body)
	if err != nil {
		return nil, err
	}

	response := new(FxRateResponse)

	_, err = f.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (f FxIntImpl) Exchange(ctx context.Context, body FxBody) (*Transaction, error) {
	req, err := f.client.NewRequest(ctx, http.MethodPost, "fx/exchange", body)
	if err != nil {
		return nil, err
	}

	response := new(Transaction)

	_, err = f.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
