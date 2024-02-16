package main

import (
	"context"
	"net/http"
)

type Business struct {
	Address   string `json:"address"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	Logo      string `json:"logo"`
	Slug      string `json:"slug"`
	Type      string `json:"type"`
	UpdatedAt string `json:"updated_at"`
}

type BusinessInt interface {
	Get(ctx context.Context) (*Business, error)
}

type BusinessIntImpl struct {
	client *SwervpayClient
}

var _ BusinessInt = &BusinessIntImpl{}

func (b *BusinessIntImpl) Get(ctx context.Context) (*Business, error) {
	// Prepare request
	req, err := b.client.NewRequest(ctx, http.MethodGet, "business", nil)
	if err != nil {
		return nil, err
	}

	response := new(Business)

	_, err = b.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
