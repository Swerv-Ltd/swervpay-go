package main

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	client *SwervpayClient
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewSwervpayClient(&SwervpayClientOption{
		BusinessID: "",
		SecretKey:  "",
		BaseURL:    "",
	})
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestGetBusiness(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/business", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &Business{
			Address: "address",
			ID:      "bus_123456789",
			Name:    "Swervpay",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Business.Get(context.Background())
	if err != nil {
		t.Errorf("Unable to get business: %v", err)
	}
	assert.Equal(t, resp.ID, "bus_123456789")
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}
