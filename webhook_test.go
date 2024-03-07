package swervpay

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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

	mux.HandleFunc("/webhook/"+logId+"/retry", func(w http.ResponseWriter, r *http.Request) {
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
