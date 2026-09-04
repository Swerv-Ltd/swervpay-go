package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvoiceGets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":"inv_001","title":"March invoice"}]`))
	})

	response, err := client.Invoice.Gets(context.Background(), &PageAndLimitQuery{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "inv_001", response[0].ID)
}

func TestInvoiceCreateGetAndUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		var body CreateInvoiceBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "cus_001", body.CustomerID)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"inv_001","title":"March invoice"}`))
	})
	mux.HandleFunc("/invoices/inv_001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"id":"inv_001","title":"March invoice"}`))
			return
		}
		testMethod(t, r, http.MethodPost)
		_, _ = w.Write([]byte(`{"id":"inv_001","title":"Updated invoice"}`))
	})

	created, err := client.Invoice.Create(context.Background(), &CreateInvoiceBody{CustomerID: "cus_001"})
	assert.NoError(t, err)
	assert.Equal(t, "inv_001", created.ID)

	fetched, err := client.Invoice.Get(context.Background(), "inv_001")
	assert.NoError(t, err)
	assert.Equal(t, "March invoice", fetched.Title)

	updated, err := client.Invoice.Update(context.Background(), "inv_001", &UpdateInvoiceBody{Title: "Updated invoice"})
	assert.NoError(t, err)
	assert.Equal(t, "Updated invoice", updated.Title)
}
