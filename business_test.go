package swervpay

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

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
