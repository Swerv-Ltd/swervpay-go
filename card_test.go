package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFundCard(t *testing.T) {
	setup()
	defer teardown()

	cardId := "card_12345678"

	mux.HandleFunc("/cards/"+cardId+"/fund", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &CardActionResponse{
			Message: "Card funded successfully",
			Transaction: &Transaction{
				ID:        "txn_card_fund_001",
				Amount:    1000,
				Status:    "success",
				Reference: "ref_001",
				Type:      "fund",
				Category:  "card",
			},
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
	if assert.NotNil(t, resp.Transaction) {
		assert.Equal(t, resp.Transaction.ID, "txn_card_fund_001")
		assert.Equal(t, resp.Transaction.Amount, 1000.0)
		assert.Equal(t, resp.Transaction.Status, "success")
	}
}

func TestWithdrawCard(t *testing.T) {
	setup()
	defer teardown()

	cardId := "card_12345678"

	mux.HandleFunc("/cards/"+cardId+"/withdraw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &CardActionResponse{
			Message: "Card withdrawal successful",
			Transaction: &Transaction{
				ID:        "txn_card_withdraw_001",
				Amount:    500,
				Status:    "success",
				Reference: "ref_002",
				Type:      "withdraw",
				Category:  "card",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &FundOrWithdrawCardBody{
		Amount: 500,
	}

	resp, err := client.Card.Withdraw(context.Background(), cardId, req)
	if err != nil {
		t.Errorf("Unable to withdraw from card: %v", err)
	}
	assert.Equal(t, resp.Message, "Card withdrawal successful")
	if assert.NotNil(t, resp.Transaction) {
		assert.Equal(t, resp.Transaction.ID, "txn_card_withdraw_001")
		assert.Equal(t, resp.Transaction.Amount, 500.0)
		assert.Equal(t, resp.Transaction.Status, "success")
	}
}

func TestSetCardPin(t *testing.T) {
	setup()
	defer teardown()

	cardID := "card_001"
	mux.HandleFunc("/cards/"+cardID+"/pin", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		var body SetCardPinBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "1234", body.PIN)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"card PIN updated"}`))
	})

	response, err := client.Card.SetPin(context.Background(), cardID, &SetCardPinBody{PIN: "1234"})
	assert.NoError(t, err)
	assert.Equal(t, "card PIN updated", response.Message)
}
