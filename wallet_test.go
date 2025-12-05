package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletGets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*Wallet{
			{
				ID:            "wallet_001",
				AccountName:   "John Doe",
				AccountNumber: "1234567890",
				Balance:       5000.0,
				BankCode:      "058",
				BankName:      "GTBank",
			},
			{
				ID:            "wallet_002",
				AccountName:   "Jane Smith",
				AccountNumber: "0987654321",
				Balance:       3000.0,
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

	resp, err := client.Wallet.Gets(context.Background(), query)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "wallet_001")
	assert.Equal(t, resp[0].AccountName, "John Doe")
	assert.Equal(t, resp[0].Balance, 5000.0)
	assert.Equal(t, resp[1].ID, "wallet_002")
	assert.Equal(t, resp[1].Balance, 3000.0)
}

func TestWalletGet(t *testing.T) {
	setup()
	defer teardown()

	walletId := "wallet_001"

	mux.HandleFunc("/wallets/"+walletId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Wallet{
			ID:            walletId,
			AccountName:   "John Doe",
			AccountNumber: "1234567890",
			Balance:       5000.0,
			BankCode:      "058",
			BankName:      "GTBank",
			IsBlocked:     false,
			Label:         "Main Wallet",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Wallet.Get(context.Background(), walletId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, walletId)
	assert.Equal(t, resp.AccountName, "John Doe")
	assert.Equal(t, resp.Balance, 5000.0)
	assert.Equal(t, resp.BankCode, "058")
	assert.False(t, resp.IsBlocked)
}

func TestWalletCredit(t *testing.T) {
	setup()
	defer teardown()

	walletId := "wallet_001"

	mux.HandleFunc("/wallets/"+walletId+"/credit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &CreditWalletResponse{
			ID:        "txn_123456",
			Message:   "Wallet credited successfully",
			Reference: "ref_123456",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CreditWalletBody{
		Amount: 1000.0,
		Sender: CreditWalletSenderInput{
			AccountName:   "Sender Name",
			AccountNumber: "9876543210",
			BankCode:      "058",
			BankName:      "GTBank",
			Narration:     "Credit transaction",
			Reference:     "ref_123456",
		},
	}

	resp, err := client.Wallet.Credit(context.Background(), walletId, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, "txn_123456")
	assert.Equal(t, resp.Message, "Wallet credited successfully")
	assert.Equal(t, resp.Reference, "ref_123456")
}
