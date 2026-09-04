package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckoutCredit(t *testing.T) {
	setup()
	defer teardown()

	accountNo := "1234567890"
	mux.HandleFunc("/checkouts/"+accountNo+"/credit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		var body CreditCheckoutBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 1000.0, body.Amount)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"txn_001","message":"credited","reference":"ref_001"}`))
	})

	response, err := client.Checkout.Credit(context.Background(), accountNo, &CreditCheckoutBody{Amount: 1000})
	assert.NoError(t, err)
	assert.Equal(t, "txn_001", response.ID)
	assert.Equal(t, "ref_001", response.Reference)
}
