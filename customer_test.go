package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerGets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*Customer{
			{
				ID:          "cust_001",
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "john.doe@example.com",
				PhoneNumber: "+2348012345678",
				Country:     "NG",
				Status:      "active",
			},
			{
				ID:          "cust_002",
				FirstName:   "Jane",
				LastName:    "Smith",
				Email:       "jane.smith@example.com",
				PhoneNumber: "+2348098765432",
				Country:     "NG",
				Status:      "active",
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

	resp, err := client.Customer.Gets(context.Background(), query)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "cust_001")
	assert.Equal(t, resp[0].FirstName, "John")
	assert.Equal(t, resp[1].ID, "cust_002")
	assert.Equal(t, resp[1].FirstName, "Jane")
}

func TestCustomerGet(t *testing.T) {
	setup()
	defer teardown()

	customerId := "cust_001"

	mux.HandleFunc("/customers/"+customerId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Customer{
			ID:            customerId,
			FirstName:     "John",
			LastName:      "Doe",
			MiddleName:    "Michael",
			Email:         "john.doe@example.com",
			PhoneNumber:   "+2348012345678",
			Country:       "NG",
			Status:        "active",
			IsBlacklisted: false,
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Customer.Get(context.Background(), customerId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, customerId)
	assert.Equal(t, resp.FirstName, "John")
	assert.Equal(t, resp.LastName, "Doe")
	assert.Equal(t, resp.Email, "john.doe@example.com")
	assert.False(t, resp.IsBlacklisted)
}

func TestCustomerCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Customer{
			ID:            "cust_001",
			FirstName:     "John",
			LastName:      "Doe",
			MiddleName:    "Michael",
			Email:         "john.doe@example.com",
			Country:       "NG",
			Status:        "active",
			IsBlacklisted: false,
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CreateCustomerBody{
		Country:    "NG",
		Email:      "john.doe@example.com",
		Firstname:  "John",
		Lastname:   "Doe",
		Middlename: "Michael",
	}

	resp, err := client.Customer.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, "cust_001")
	assert.Equal(t, resp.FirstName, "John")
	assert.Equal(t, resp.LastName, "Doe")
	assert.Equal(t, resp.Email, "john.doe@example.com")
}

func TestCustomerUpdate(t *testing.T) {
	setup()
	defer teardown()

	customerId := "cust_001"

	mux.HandleFunc("/customers/"+customerId+"/update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Customer{
			ID:          customerId,
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "john.updated@example.com",
			PhoneNumber: "+2348012345678",
			Country:     "NG",
			Status:      "active",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &UpdateustomerBody{
		Email:       "john.updated@example.com",
		PhoneNumber: "+2348012345678",
	}

	resp, err := client.Customer.Update(context.Background(), customerId, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.ID, customerId)
	assert.Equal(t, resp.Email, "john.updated@example.com")
	assert.Equal(t, resp.PhoneNumber, "+2348012345678")
}

func TestCustomerKyc(t *testing.T) {
	setup()
	defer teardown()

	customerId := "cust_001"

	mux.HandleFunc("/customers/"+customerId+"/kyc", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &DefaultResponse{
			Message: "KYC information updated successfully",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CustomerKycBody{
		Tier: "1",
		Tier1: Tier1KycInput{
			Bvn:         "12345678901",
			State:       "Lagos",
			City:        "Ikeja",
			Country:     "NG",
			Address:     "123 Main Street",
			PostalCode:  "100001",
			PhoneNumber: "+2348012345678",
		},
	}

	resp, err := client.Customer.Kyc(context.Background(), customerId, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Message, "KYC information updated successfully")
}

func TestCustomerBlacklist(t *testing.T) {
	setup()
	defer teardown()

	customerId := "cust_001"

	mux.HandleFunc("/customers/"+customerId+"/blacklist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &DefaultResponse{
			Message: "Customer blacklisted successfully",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Customer.Blacklist(context.Background(), customerId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Message, "Customer blacklisted successfully")
}
