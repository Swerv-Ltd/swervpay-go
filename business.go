package main

import (
	"context"
	"net/http"
)

type Business struct {
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
