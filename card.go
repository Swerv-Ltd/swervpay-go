package main

import (
	"context"
	"net/http"
)

type Card struct {
	AddressCity       string `json:"address_city"`
	AddressCountry    string `json:"address_country"`
	AddressPostalCode string `json:"address_postal_code"`
	AddressState      string `json:"address_state"`
	AddressStreet     string `json:"address_street"`
	Balance           int64  `json:"balance"`
	BusinessID        string `json:"business_id"`
	CardNumber        string `json:"card_number"`
	CreatedAt         string `json:"created_at"`
	Currency          string `json:"currency"`
	CustomerID        string `json:"customer_id"`
	Cvv               string `json:"cvv"`
	Expiry            string `json:"expiry"`
	Freeze            bool   `json:"freeze"`
	ID                string `json:"id"`
	Issuer            string `json:"issuer"`
	MaskedPan         string `json:"masked_pan"`
	NameOnCard        string `json:"name_on_card"`
	Status            string `json:"status"`
	TotalFunded       int64  `json:"total_funded"`
	Type              string `json:"type"`
	UpdatedAt         string `json:"updated_at"`
}

type CreateCardBody struct {
}

type FundOrWithdrawCardBody struct {
	Amount float64 `json:"amount"`
}

type CardInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error)
	Get(ctx context.Context, id string) (*Card, error)
	Create(ctx context.Context, body CreateCardBody) (*DefaultResponse, error)
	Fund(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error)
	Withdraw(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error)
	Terminate(ctx context.Context, id string) (*DefaultResponse, error)
	Freeze(ctx context.Context, id string) (*DefaultResponse, error)
}

type CardIntImpl struct {
	client *SwervpayClient
}

var _ CardInt = &CardIntImpl{}

func (c CardIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error) {
	path := GenerateURLPath("cards", query)

	// Prepare request
	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Card)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Get(ctx context.Context, id string) (*Card, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "cards/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Card)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Create(ctx context.Context, body CreateCardBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards", body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Fund(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/fund", body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Withdraw(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/withdraw", body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Terminate(ctx context.Context, id string) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/terminate", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c CardIntImpl) Freeze(ctx context.Context, id string) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/freeze", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
