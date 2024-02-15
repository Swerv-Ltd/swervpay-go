package main

import (
	"context"
	"net/http"
)

type WebhookInt interface {
	Test(ctx context.Context, id string) (*DefaultResponse, error)
	Retry(ctx context.Context, logId string) (*DefaultResponse, error)
}

type WebhookIntImpl struct {
	client *SwervpayClient
}

var _ WebhookInt = &WebhookIntImpl{}

func (w WebhookIntImpl) Test(ctx context.Context, id string) (*DefaultResponse, error) {
	// Prepare request
	req, err := w.client.NewRequest(ctx, http.MethodPost, "webhook/"+id+"/test", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (w WebhookIntImpl) Retry(ctx context.Context, logId string) (*DefaultResponse, error) {
	// Prepare request
	req, err := w.client.NewRequest(ctx, http.MethodPost, "webhook/"+logId+"/retry", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = w.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
