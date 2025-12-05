package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBanks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/banks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*Bank{
			{
				Code: "058",
				Name: "GTBank",
			},
			{
				Code: "044",
				Name: "Access Bank",
			},
			{
				Code: "050",
				Name: "Ecobank",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Other.Banks(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 3)
	assert.Equal(t, resp[0].Code, "058")
	assert.Equal(t, resp[0].Name, "GTBank")
	assert.Equal(t, resp[1].Code, "044")
	assert.Equal(t, resp[1].Name, "Access Bank")
	assert.Equal(t, resp[2].Code, "050")
	assert.Equal(t, resp[2].Name, "Ecobank")
}

func TestResolveAccountNumber(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/resolve-account-number", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &ResolveAccountNumber{
			AccountNumber: "1234567890",
			BankCode:      "058",
			BankName:      "GTBank",
			AccountName:   "John Doe",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := ResolveAccountNumberBody{
		AccountNumber: "1234567890",
		BankCode:      "058",
	}

	resp, err := client.Other.ResolveAccountNumber(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.AccountNumber, "1234567890")
	assert.Equal(t, resp.BankCode, "058")
	assert.Equal(t, resp.BankName, "GTBank")
	assert.Equal(t, resp.AccountName, "John Doe")
}
