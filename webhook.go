package swervpay

import (
	"context"
	"net/http"
)

// WebhookInt is an interface that defines two methods: Test and Retry.
type WebhookInt interface {
	// Test sends a test webhook request.
	// It takes a context and an id as parameters.
	// It returns a pointer to a DefaultResponse and an error.
	Test(ctx context.Context, id string) (*DefaultResponse, error)

	// Retry retries a failed webhook request.
	// It takes a context and a logId as parameters.
	// It returns a pointer to a DefaultResponse and an error.
	Retry(ctx context.Context, logId string) (*DefaultResponse, error)
}

// WebhookIntImpl is a struct that implements the WebhookInt interface.
// It contains a client of type *SwervpayClient.
type WebhookIntImpl struct {
	client *SwervpayClient
}

// Assert that WebhookIntImpl implements the WebhookInt interface.
var _ WebhookInt = &WebhookIntImpl{}

// Test is a method of WebhookIntImpl that sends a test webhook request.
// It prepares a new request and performs it.
// If there is an error during these operations, it returns the error.
// Otherwise, it returns the response and nil.
// https://docs.swervpay.co/api-reference/webhook/test
func (w WebhookIntImpl) Test(ctx context.Context, id string) (*DefaultResponse, error) {
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

// Retry is a method of WebhookIntImpl that retries a failed webhook request.
// It prepares a new request and performs it.
// If there is an error during these operations, it returns the error.
// Otherwise, it returns the response and nil.
// https://docs.swervpay.co/api-reference/webhook/retry
func (w WebhookIntImpl) Retry(ctx context.Context, logId string) (*DefaultResponse, error) {
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
