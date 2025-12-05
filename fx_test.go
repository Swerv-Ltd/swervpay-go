package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFxRate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/fx/rate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &FxRateResponse{
			Rate: 1500.0,
			From: FromOrTo{
				Amount:   "1000",
				Currency: "NGN",
			},
			To: FromOrTo{
				Amount:   "0.67",
				Currency: "USD",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := FxBody{
		Amount: 1000.0,
		From:   "NGN",
		To:     "USD",
	}

	resp, err := client.Fx.Rate(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Rate, 1500.0)
	assert.Equal(t, resp.From.Currency, "NGN")
	assert.Equal(t, resp.From.Amount, "1000")
	assert.Equal(t, resp.To.Currency, "USD")
	assert.Equal(t, resp.To.Amount, "0.67")
}

func TestFxExchange(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/fx/exchange", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Transaction{
			ID:            "fx_txn_001",
			Amount:        0.67,
			Status:        "success",
			Reference:     "fx_ref_001",
			Type:          "fx",
			Category:      "exchange",
			Charges:       5.0,
			FiatRate:      1500.0,
			AccountName:   "John Doe",
			AccountNumber: "1234567890",
			BankCode:      "058",
			BankName:      "GTBank",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := FxBody{
		Amount: 1000.0,
		From:   "NGN",
		To:     "USD",
	}

	resp, err := client.Fx.Exchange(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, "fx_txn_001")
	assert.Equal(t, resp.Amount, 0.67)
	assert.Equal(t, resp.Status, "success")
	assert.Equal(t, resp.Type, "fx")
	assert.Equal(t, resp.FiatRate, 1500.0)
}
