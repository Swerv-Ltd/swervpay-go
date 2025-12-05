package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionGets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*Transaction{
			{
				ID:            "txn_001",
				Amount:        1000.0,
				Status:        "success",
				Reference:     "ref_001",
				Type:          "credit",
				Category:      "transfer",
				Charges:       10.0,
				AccountName:   "John Doe",
				AccountNumber: "1234567890",
				BankCode:      "058",
				BankName:      "GTBank",
			},
			{
				ID:            "txn_002",
				Amount:        500.0,
				Status:        "success",
				Reference:     "ref_002",
				Type:          "debit",
				Category:      "payout",
				Charges:       5.0,
				AccountName:   "Jane Smith",
				AccountNumber: "0987654321",
				BankCode:      "044",
				BankName:      "Access Bank",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	query := &PageAndLimitQuery{
		Page:  1,
		Limit: 10,
	}

	resp, err := client.Transaction.Gets(context.Background(), query)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "txn_001")
	assert.Equal(t, resp[0].Amount, 1000.0)
	assert.Equal(t, resp[0].Status, "success")
	assert.Equal(t, resp[1].ID, "txn_002")
	assert.Equal(t, resp[1].Amount, 500.0)
}

func TestTransactionGet(t *testing.T) {
	setup()
	defer teardown()

	transactionId := "txn_001"

	mux.HandleFunc("/transactions/"+transactionId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Transaction{
			ID:            transactionId,
			Amount:        1000.0,
			Status:        "success",
			Reference:     "ref_001",
			Type:          "credit",
			Category:      "transfer",
			Charges:       10.0,
			AccountName:   "John Doe",
			AccountNumber: "1234567890",
			BankCode:      "058",
			BankName:      "GTBank",
			CreatedAt:     "2024-01-01T00:00:00Z",
			UpdatedAt:     "2024-01-01T00:00:00Z",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Transaction.Get(context.Background(), transactionId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, transactionId)
	assert.Equal(t, resp.Amount, 1000.0)
	assert.Equal(t, resp.Status, "success")
	assert.Equal(t, resp.Type, "credit")
	assert.Equal(t, resp.Charges, 10.0)
}
