package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionGets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*Wallet{
			{
				ID:            "coll_001",
				AccountName:   "Collection Account 1",
				AccountNumber: "1234567890",
				Balance:       10000.0,
				BankCode:      "058",
				BankName:      "GTBank",
			},
			{
				ID:            "coll_002",
				AccountName:   "Collection Account 2",
				AccountNumber: "0987654321",
				Balance:       5000.0,
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

	resp, err := client.Collection.Gets(context.Background(), query)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "coll_001")
	assert.Equal(t, resp[0].Balance, 10000.0)
	assert.Equal(t, resp[1].ID, "coll_002")
	assert.Equal(t, resp[1].Balance, 5000.0)
}

func TestCollectionGet(t *testing.T) {
	setup()
	defer teardown()

	collectionId := "coll_001"

	mux.HandleFunc("/collections/"+collectionId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Wallet{
			ID:            collectionId,
			AccountName:   "Collection Account 1",
			AccountNumber: "1234567890",
			Balance:       10000.0,
			BankCode:      "058",
			BankName:      "GTBank",
			Reference:     "ref_001",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Collection.Get(context.Background(), collectionId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, collectionId)
	assert.Equal(t, resp.AccountName, "Collection Account 1")
	assert.Equal(t, resp.Balance, 10000.0)
	assert.Equal(t, resp.BankCode, "058")
}

func TestCollectionCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Wallet{
			ID:            "coll_001",
			AccountName:   "Collection Account",
			AccountNumber: "1234567890",
			Balance:       0.0,
			BankCode:      "058",
			BankName:      "GTBank",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CreateCollectionBody{
		CustomerID:   "cust_001",
		Currency:     "NGN",
		MerchantName: "Test Merchant",
		Amount:       1000.0,
		Type:         "collection",
	}

	resp, err := client.Collection.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, "coll_001")
	assert.Equal(t, resp.AccountName, "Collection Account")
	assert.Equal(t, resp.BankCode, "058")
}

func TestCollectionCredit(t *testing.T) {
	setup()
	defer teardown()

	collectionId := "coll_001"

	mux.HandleFunc("/collections/"+collectionId+"/credit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &CreditWalletResponse{
			ID:        "txn_123456",
			Message:   "Collection credited successfully",
			Reference: "ref_123456",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CreditWalletBody{
		Amount: 2000.0,
		Sender: CreditWalletSenderInput{
			AccountName:   "Sender Name",
			AccountNumber: "9876543210",
			BankCode:      "058",
			BankName:      "GTBank",
			Narration:     "Credit collection",
			Reference:     "ref_123456",
		},
	}

	resp, err := client.Collection.Credit(context.Background(), collectionId, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, "txn_123456")
	assert.Equal(t, resp.Message, "Collection credited successfully")
	assert.Equal(t, resp.Reference, "ref_123456")
}

func TestCollectionTransactions(t *testing.T) {
	setup()
	defer teardown()

	collectionId := "coll_001"

	mux.HandleFunc("/collections/"+collectionId+"/transactions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*CollectionHistory{
			{
				ID:            "hist_001",
				Amount:        1000.0,
				Charges:       10.0,
				Currency:      "NGN",
				PaymentMethod: "card",
				Reference:     "ref_001",
				CreatedAt:     "2024-01-01T00:00:00Z",
				UpdatedAt:     "2024-01-01T00:00:00Z",
			},
			{
				ID:            "hist_002",
				Amount:        500.0,
				Charges:       5.0,
				Currency:      "NGN",
				PaymentMethod: "bank_transfer",
				Reference:     "ref_002",
				CreatedAt:     "2024-01-02T00:00:00Z",
				UpdatedAt:     "2024-01-02T00:00:00Z",
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

	resp, err := client.Collection.Transactions(context.Background(), collectionId, query)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "hist_001")
	assert.Equal(t, resp[0].Amount, 1000.0)
	assert.Equal(t, resp[0].Charges, 10.0)
	assert.Equal(t, resp[1].ID, "hist_002")
	assert.Equal(t, resp[1].Amount, 500.0)
}
