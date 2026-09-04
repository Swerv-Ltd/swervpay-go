package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityBVN(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/identity/bvn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		var body BVNRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "12345678901", body.Number)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"bvn":"12345678901","first_name":"Ada","lga_of_Origin":"Ikeja"}`))
	})

	response, err := client.Identity.BVN(context.Background(), &BVNRequest{Number: "12345678901"})
	assert.NoError(t, err)
	assert.Equal(t, "12345678901", response.BVN)
	assert.Equal(t, "Ada", response.FirstName)
	assert.Equal(t, "Ikeja", response.LGAOfOrigin)
}
