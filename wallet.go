package main

import (
	"context"
	"net/http"
)

type Wallet struct {
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
