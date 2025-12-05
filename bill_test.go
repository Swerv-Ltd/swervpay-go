package swervpay

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBill(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bills", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &BillTransaction{
			ID:            "bill_123456789",
			Amount:        1000.0,
			Status:        "success",
			Reference:     "ref_123456",
			Category:      "electricity",
			PaymentMethod: "card",
			AccountName:   "John Doe",
			AccountNumber: "1234567890",
			BankCode:      "058",
			BankName:      "GTBank",
			Charges:       50.0,
			Bill: &BillDetail{
				BillCode: "ELEC001",
				BillName: "Electricity Bill",
				ItemCode: "ITEM001",
				Name:     "Monthly Electricity",
				Token:    "token_123456",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &CreateBillBody{
		Amount:     1000.0,
		BillerID:   "biller_123",
		Category:   "electricity",
		CustomerID: "cust_123456",
		ItemID:     "item_123",
		Reference:  "ref_123456",
	}

	resp, err := client.Bill.Create(context.Background(), req)
	if err != nil {
		t.Errorf("Unable to create bill: %v", err)
	}
	assert.Equal(t, resp.ID, "bill_123456789")
	assert.Equal(t, resp.Amount, 1000.0)
	assert.Equal(t, resp.Status, "success")
	assert.NotNil(t, resp.Bill)
	assert.Equal(t, resp.Bill.BillCode, "ELEC001")
}

func TestGetBill(t *testing.T) {
	setup()
	defer teardown()

	billId := "bill_123456789"

	mux.HandleFunc("/bills/"+billId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &BillTransaction{
			ID:            billId,
			Amount:        1000.0,
			Status:        "success",
			Reference:     "ref_123456",
			Category:      "electricity",
			PaymentMethod: "card",
			Charges:       50.0,
			Bill: &BillDetail{
				BillCode: "ELEC001",
				BillName: "Electricity Bill",
				ItemCode: "ITEM001",
				Name:     "Monthly Electricity",
				Token:    "token_123456",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Bill.Get(context.Background(), billId)
	if err != nil {
		t.Errorf("Unable to get bill: %v", err)
	}
	assert.Equal(t, resp.ID, billId)
	assert.Equal(t, resp.Amount, 1000.0)
	assert.Equal(t, resp.Status, "success")
	assert.Equal(t, resp.Charges, 50.0)
	assert.NotNil(t, resp.Bill)
}

func TestBillCategories(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bills/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		// Verify query parameters
		assert.Equal(t, r.URL.Query().Get("page"), "1")
		assert.Equal(t, r.URL.Query().Get("limit"), "10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*BillCategory{
			{
				ID:   "cat_001",
				Name: "Electricity",
			},
			{
				ID:   "cat_002",
				Name: "Water",
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

	resp, err := client.Bill.Categories(context.Background(), query)
	if err != nil {
		t.Errorf("Unable to get bill categories: %v", err)
		return
	}
	if len(resp) == 0 {
		t.Errorf("Expected 2 categories, got empty response")
		return
	}
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "cat_001")
	assert.Equal(t, resp[0].Name, "Electricity")
	assert.Equal(t, resp[1].ID, "cat_002")
	assert.Equal(t, resp[1].Name, "Water")
}

func TestBillCategoryLists(t *testing.T) {
	setup()
	defer teardown()

	categoryId := "cat_001"

	mux.HandleFunc("/bills/categories/"+categoryId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*BillerList{
			{
				ID:   "biller_001",
				Name: "Eko Electricity Distribution Company",
			},
			{
				ID:   "biller_002",
				Name: "Ikeja Electric",
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Bill.CategoryLists(context.Background(), categoryId)
	if err != nil {
		t.Errorf("Unable to get bill category lists: %v", err)
	}
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "biller_001")
	assert.Equal(t, resp[0].Name, "Eko Electricity Distribution Company")
	assert.Equal(t, resp[1].ID, "biller_002")
	assert.Equal(t, resp[1].Name, "Ikeja Electric")
}

func TestBillCategoryListItems(t *testing.T) {
	setup()
	defer teardown()

	categoryId := "cat_001"
	itemId := "biller_001"

	mux.HandleFunc("/bills/categories/"+categoryId+"/items/"+itemId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := []*BillerItem{
			{
				ID:       "item_001",
				Name:     "Prepaid Meter",
				Code:     "PREPAID001",
				Amount:   5000.0,
				Currency: "NGN",
				Fee:      50.0,
			},
			{
				ID:       "item_002",
				Name:     "Postpaid Meter",
				Code:     "POSTPAID001",
				Amount:   3000.0,
				Currency: "NGN",
				Fee:      30.0,
			},
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	resp, err := client.Bill.CategoryListItems(context.Background(), categoryId, itemId)
	if err != nil {
		t.Errorf("Unable to get bill category list items: %v", err)
	}
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, "item_001")
	assert.Equal(t, resp[0].Name, "Prepaid Meter")
	assert.Equal(t, resp[0].Amount, 5000.0)
	assert.Equal(t, resp[0].Fee, 50.0)
	assert.Equal(t, resp[1].ID, "item_002")
	assert.Equal(t, resp[1].Amount, 3000.0)
}

func TestValidateBill(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bills/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Validate endpoint returns no response body
	})

	req := &ValidateBillBody{
		BillerID:   "biller_123",
		Category:   "electricity",
		CustomerID: "cust_123456",
		ItemID:     "item_123",
	}

	err := client.Bill.Validate(context.Background(), req)
	if err != nil {
		t.Errorf("Unable to validate bill: %v", err)
	}
	// Validate endpoint returns no error on success
	assert.NoError(t, err)
}
