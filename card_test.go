package swervpay

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFundCard(t *testing.T) {
	setup()
	defer teardown()

	cardId := "card_12345678"

	mux.HandleFunc("/cards/"+cardId+"/fund", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &DefaultResponse{
			Message: "Card funded successfully",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &FundOrWithdrawCardBody{
		Amount: 1000,
	}

	resp, err := client.Card.Fund(context.Background(), cardId, req)
	if err != nil {
		t.Errorf("Unable to fund card: %v", err)
	}
	assert.Equal(t, resp.Message, "Card funded successfully")
}
