package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestWebhook(t *testing.T) {
	setup()
	defer teardown()

	webhookId := "wbh_123456789"

	mux.HandleFunc("/webhook/"+webhookId+"/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &DefaultResponse{
			Message: "Test webhook sent successfully",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Webhook.Test(context.Background(), webhookId)
	if err != nil {
		t.Errorf("Err retrying webhook: %v", err)
	}
	assert.Equal(t, resp.Message, "Test webhook sent successfully")
}

func TestRetryWebhook(t *testing.T) {
	setup()
	defer teardown()

	logId := "tri_123456789"

	mux.HandleFunc("/webhook/log/"+logId+"/retry", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &DefaultResponse{
			Message: "Retry webhook sent successfully",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Webhook.Retry(context.Background(), logId)
	if err != nil {
		t.Errorf("Err retrying webhook: %v", err)
	}
	assert.Equal(t, resp.Message, "Retry webhook sent successfully")
}

func TestCreateWebhook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		var body CreateWebhookBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "https://example.com/webhook", body.URL)
		assert.Equal(t, WebhookEventCustomerCreated, body.Events[0])
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"wh_001","endpoint_url":"https://example.com/webhook"}`))
	})

	response, err := client.Webhook.Create(context.Background(), &CreateWebhookBody{
		URL:         "https://example.com/webhook",
		Events:      []WebhookEvent{WebhookEventCustomerCreated},
		Description: "customer events",
	})
	assert.NoError(t, err)
	assert.Equal(t, "wh_001", response.ID)
	assert.Equal(t, "https://example.com/webhook", response.EndpointURL)
}

func TestWebhookLogs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/webhook/logs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "failed", r.URL.Query().Get("status"))
		assert.Equal(t, "wh_001", r.URL.Query().Get("webhook_id"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":"log_001","event":"customer.created","status":"failed"}]`))
	})

	response, err := client.Webhook.Logs(context.Background(), &WebhookLogsQuery{
		Page:      1,
		Limit:     10,
		Status:    "failed",
		WebhookID: "wh_001",
	})
	assert.NoError(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, WebhookEventCustomerCreated, response[0].Event)
}

func TestEditWebhook(t *testing.T) {
	setup()
	defer teardown()

	webhookID := "wh_001"
	mux.HandleFunc("/webhook/"+webhookID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		var body EditWebhookBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "https://example.com/updated", body.URL)
		_, _ = w.Write([]byte(`{"message":"webhook updated"}`))
	})

	response, err := client.Webhook.Edit(context.Background(), webhookID, &EditWebhookBody{
		URL: "https://example.com/updated",
	})
	assert.NoError(t, err)
	assert.Equal(t, "webhook updated", response.Message)
}
