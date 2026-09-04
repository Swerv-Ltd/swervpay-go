package swervpay

import (
	"context"
	"net/http"
)

// WebhookEvent identifies an event that can trigger a webhook.
type WebhookEvent string

// WebhookEventTrigger is an alias for WebhookEvent.
type WebhookEventTrigger = WebhookEvent

// Webhook event names supported by the API.
const (
	WebhookEventBillCompleted             WebhookEvent = "bill.completed"
	WebhookEventBillFailed                WebhookEvent = "bill.failed"
	WebhookEventBillProcessed             WebhookEvent = "bill.processed"
	WebhookEventCardAuthorization         WebhookEvent = "card.authorization"
	WebhookEventCardCharges               WebhookEvent = "card.charges"
	WebhookEventCardContactlessActivation WebhookEvent = "card.contactless.activation"
	WebhookEventCardCreated               WebhookEvent = "card.created"
	WebhookEventCardCreatedFailed         WebhookEvent = "card.created.failed"
	WebhookEventCardFreezed               WebhookEvent = "card.freezed"
	WebhookEventCardTerminated            WebhookEvent = "card.terminated"
	WebhookEventCardTransaction           WebhookEvent = "card.transaction"
	WebhookEventCardUnfreezed             WebhookEvent = "card.unfreezed"
	WebhookEventCardUpdated               WebhookEvent = "card.updated"
	WebhookEventCheckoutCompleted         WebhookEvent = "checkout.completed"
	WebhookEventCheckoutFailed            WebhookEvent = "checkout.failed"
	WebhookEventCollectionCompleted       WebhookEvent = "collection.completed"
	WebhookEventCollectionCreated         WebhookEvent = "collection.created"
	WebhookEventCollectionCreatedFailed   WebhookEvent = "collection.created.failed"
	WebhookEventCollectionFailed          WebhookEvent = "collection.failed"
	WebhookEventCollectionNotification    WebhookEvent = "collection.notification"
	WebhookEventCollectionUpdated         WebhookEvent = "collection.updated"
	WebhookEventCustomerCreated           WebhookEvent = "customer.created"
	WebhookEventCustomerKYCUpdated        WebhookEvent = "customer.kyc.updated"
	WebhookEventCustomerUpdated           WebhookEvent = "customer.updated"
	WebhookEventPayoutCompleted           WebhookEvent = "payout.completed"
	WebhookEventPayoutFailed              WebhookEvent = "payout.failed"
	WebhookEventPayoutProcessed           WebhookEvent = "payout.processed"
	WebhookEventPayoutProcessing          WebhookEvent = "payout.processing"
	WebhookEventPayoutReversed            WebhookEvent = "payout.reversed"
	WebhookEventWalletCreated             WebhookEvent = "wallet.created"
	WebhookEventWalletUpdated             WebhookEvent = "wallet.updated"
)

// Webhook represents a configured webhook.
type Webhook struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	EndpointURL string `json:"endpoint_url"`
	Events      string `json:"events"`
	ID          string `json:"id"`
	SigningKey  string `json:"signing_key"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updated_at"`
}

// WebhookLog represents a webhook delivery log.
type WebhookLog struct {
	BusinessWebhookID  string       `json:"business_webhook_id"`
	CreatedAt          string       `json:"created_at"`
	Event              WebhookEvent `json:"event"`
	ID                 string       `json:"id"`
	RequestBody        string       `json:"request_body"`
	ResponseBody       string       `json:"response_body"`
	ResponseStatusCode string       `json:"response_status_code"`
	Status             string       `json:"status"`
	UpdatedAt          string       `json:"updated_at"`
}

// CreateWebhookBody represents a webhook creation request.
type CreateWebhookBody struct {
	Description string         `json:"description"`
	Events      []WebhookEvent `json:"events"`
	URL         string         `json:"url"`
}

// EditWebhookBody represents a webhook update request.
type EditWebhookBody = CreateWebhookBody

// WebhookLogsQuery contains optional webhook log filters.
type WebhookLogsQuery struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Status    string `json:"status,omitempty"`
	Event     string `json:"event,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
}

// WebhookLogQuery is an alias for WebhookLogsQuery.
type WebhookLogQuery = WebhookLogsQuery

// WebhookInt defines webhook operations.
type WebhookInt interface {
	Create(ctx context.Context, body *CreateWebhookBody) (*Webhook, error)
	Edit(ctx context.Context, id string, body *EditWebhookBody) (*DefaultResponse, error)
	Logs(ctx context.Context, query *WebhookLogsQuery) ([]*WebhookLog, error)

	// Test sends a test webhook request.
	Test(ctx context.Context, id string) (*DefaultResponse, error)

	// Retry retries a failed webhook request.
	Retry(ctx context.Context, logId string) (*DefaultResponse, error)
}

// WebhookIntImpl is a struct that implements the WebhookInt interface.
// It contains a client of type *SwervpayClient.
type WebhookIntImpl struct {
	client *SwervpayClient
}

// Assert that WebhookIntImpl implements the WebhookInt interface.
var _ WebhookInt = &WebhookIntImpl{}

// Create creates a webhook.
func (w WebhookIntImpl) Create(ctx context.Context, body *CreateWebhookBody) (*Webhook, error) {
	req, err := w.client.NewRequest(ctx, http.MethodPost, "webhook", body)
	if err != nil {
		return nil, err
	}

	response := new(Webhook)
	if _, err = w.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Edit updates a webhook by ID.
func (w WebhookIntImpl) Edit(ctx context.Context, id string, body *EditWebhookBody) (*DefaultResponse, error) {
	req, err := w.client.NewRequest(ctx, http.MethodPut, "webhook/"+id, body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)
	if _, err = w.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Logs lists webhook delivery logs.
func (w WebhookIntImpl) Logs(ctx context.Context, query *WebhookLogsQuery) ([]*WebhookLog, error) {
	path := GenerateURLPath("webhook/logs", query)
	req, err := w.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*WebhookLog{}
	if _, err = w.client.Perform(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

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
	req, err := w.client.NewRequest(ctx, http.MethodPost, "webhook/log/"+logId+"/retry", nil)
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
