package swervpay

import (
	"context"
	"net/http"
)

type CollectionHistory struct {
	Amount        int64  `json:"amount"`
	Charges       int64  `json:"charges"`
	CreatedAt     string `json:"created_at"`
	Currency      string `json:"currency"`
	ID            string `json:"id"`
	PaymentMethod string `json:"payment_method"`
	Reference     string `json:"reference"`
	UpdatedAt     string `json:"updated_at"`
}

type CreateCollectionBody struct {
	CustomerID   string `json:"customer_id"`
	Currency     string `json:"currency"`
	MerchantName string `json:"merchant_name"`
	Amount       int64  `json:"amount"`
	Type         string `json:"type"`
}

type CollectionInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Wallet, error)
	Get(ctx context.Context, id string) (*Wallet, error)
	Create(ctx context.Context, body *CreateCollectionBody) (*Wallet, error)
	Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CollectionHistory, error)
}

type CollectionIntImpl struct {
	client *SwervpayClient
}

var _ CollectionInt = &CollectionIntImpl{}

func (c CollectionIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Wallet, error) {
	path := GenerateURLPath("collections", query)

	// Prepare request
	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Wallet)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CollectionIntImpl) Get(ctx context.Context, id string) (*Wallet, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CollectionIntImpl) Create(ctx context.Context, body *CreateCollectionBody) (*Wallet, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "collections", body)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CollectionIntImpl) Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CollectionHistory, error) {
	path := GenerateURLPath("collections/"+id+"/transactions", query)

	// Prepare request
	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]CollectionHistory)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
