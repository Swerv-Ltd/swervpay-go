package swervpay

import (
	"context"
	"net/http"
)

type Transaction struct{}

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
