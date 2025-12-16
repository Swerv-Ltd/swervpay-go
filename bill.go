package swervpay

import (
	"context"
	"net/http"
)

// BillCategory represents a bill category entry.
type BillCategory struct {
	ID   string `json:"id"`   // Category identifier.
	Name string `json:"name"` // Category name.
}

// BillerList represents a biller under a category.
type BillerList struct {
	ID   string `json:"id"`   // Biller identifier.
	Name string `json:"name"` // Biller name.
}

// BillerItem represents a billable item for a biller.
type BillerItem struct {
	Amount   float64 `json:"amount"`   // Item amount.
	Code     string  `json:"code"`     // Item code.
	Currency string  `json:"currency"` // Currency for the item.
	Fee      float64 `json:"fee"`      // Associated fee.
	ID       string  `json:"id"`       // Item identifier.
	Name     string  `json:"name"`     // Item name.
}

// CreateBillBody represents the payload to create a bill.
type CreateBillBody struct {
	Amount     float64 `json:"amount"`      // Bill amount.
	BillerID   string  `json:"biller_id"`   // Biller identifier.
	Category   string  `json:"category"`    // Category identifier.
	CustomerID string  `json:"customer_id"` // Customer identifier.
	ItemID     string  `json:"item_id"`     // Item identifier.
	Reference  string  `json:"reference"`   // Reference for idempotency.
}

// ValidateBillBody represents the payload to validate a bill for a customer.
type ValidateBillBody struct {
	BillerID   string `json:"biller_id"`   // Biller identifier.
	Category   string `json:"category"`    // Category identifier.
	CustomerID string `json:"customer_id"` // Customer identifier.
	ItemID     string `json:"item_id"`     // Item identifier.
}

// BillDetail represents the bill details nested in a transaction.
type BillDetail struct {
	BillCode string `json:"bill_code"`       // Bill code.
	BillName string `json:"bill_name"`       // Bill name.
	ItemCode string `json:"item_code"`       // Item code.
	Name     string `json:"name,omitempty"`  // Item name.
	Token    string `json:"token,omitempty"` // Bill token.
}

// BillTransaction represents a bill transaction.
type BillTransaction struct {
	AccountName   string      `json:"account_name"`             // Account holder name.
	AccountNumber string      `json:"account_number"`           // Account number.
	Amount        float64     `json:"amount"`                   // Transaction amount.
	BankCode      string      `json:"bank_code"`                // Bank code.
	BankName      string      `json:"bank_name"`                // Bank name.
	Bill          *BillDetail `json:"bill,omitempty"`           // Bill detail payload.
	Category      string      `json:"category"`                 // Category identifier.
	Charges       float64     `json:"charges"`                  // Transaction charges.
	CreatedAt     string      `json:"created_at"`               // Creation timestamp.
	Detail        string      `json:"detail"`                   // Transaction detail.
	FiatRate      float64     `json:"fiat_rate"`                // Fiat conversion rate.
	ID            string      `json:"id"`                       // Transaction ID.
	Imad          string      `json:"imad,omitempty"`           // IMAD reference.
	PaymentMethod string      `json:"payment_method,omitempty"` // Payment method.
	Reference     string      `json:"reference"`                // Reference string.
	Report        bool        `json:"report"`                   // Report status.
	ReportMessage string      `json:"report_message,omitempty"` // Report message.
	SessionID     string      `json:"session_id,omitempty"`     // Session ID.
	Status        string      `json:"status"`                   // Transaction status.
	TraceNumber   string      `json:"trace_number,omitempty"`   // Trace number.
	Type          string      `json:"type"`                     // Transaction type.
	UpdatedAt     string      `json:"updated_at"`               // Last update timestamp.
}

// CreateBillResponse represents the response from creating a bill.
type CreateBillResponse struct {
	Message     string          `json:"message"`     // Response message.
	Transaction BillTransaction `json:"transaction"` // Transaction details.
}

// BillInt defines bill-related operations.
type BillInt interface {
	Create(ctx context.Context, body *CreateBillBody) (*CreateBillResponse, error)     // Create a new bill.
	Get(ctx context.Context, id string) (*BillTransaction, error)                      // Retrieve a bill by ID.
	Categories(ctx context.Context, query *PageAndLimitQuery) ([]*BillCategory, error) // List bill categories.
	CategoryLists(ctx context.Context, id string) ([]*BillerList, error)               // List billers for a category.
	CategoryListItems(ctx context.Context, id, itemId string) ([]*BillerItem, error)   // List biller items.
	Validate(ctx context.Context, body *ValidateBillBody) error                        // Validate a bill with customer details.
}

// BillIntImpl implements BillInt.
type BillIntImpl struct {
	client *SwervpayClient // Swervpay client.
}

// Ensure BillIntImpl satisfies BillInt.
var _ BillInt = &BillIntImpl{}

// Create creates a bill.
// https://docs.swervpay.co/api-reference/bills/create
func (b BillIntImpl) Create(ctx context.Context, body *CreateBillBody) (*CreateBillResponse, error) {
	req, err := b.client.NewRequest(ctx, http.MethodPost, "bills", body)
	if err != nil {
		return nil, err
	}

	response := new(CreateBillResponse)

	_, err = b.client.Perform(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get retrieves a bill by its ID.
// https://docs.swervpay.co/api-reference/bills/get
func (b BillIntImpl) Get(ctx context.Context, id string) (*BillTransaction, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, "bills/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(BillTransaction)

	_, err = b.client.Perform(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Categories lists bill categories.
// https://docs.swervpay.co/api-reference/bills/categories
func (b BillIntImpl) Categories(ctx context.Context, query *PageAndLimitQuery) ([]*BillCategory, error) {
	path := GenerateURLPath("bills/categories", query)

	req, err := b.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*BillCategory{}

	_, err = b.client.Perform(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CategoryLists lists billers for a category.
// https://docs.swervpay.co/api-reference/bills/category-list
func (b BillIntImpl) CategoryLists(ctx context.Context, id string) ([]*BillerList, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, "bills/categories/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := []*BillerList{}

	_, err = b.client.Perform(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CategoryListItems lists biller items for a category.
// https://docs.swervpay.co/api-reference/bills/category-list-items
func (b BillIntImpl) CategoryListItems(ctx context.Context, id, itemId string) ([]*BillerItem, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, "bills/categories/"+id+"/items/"+itemId, nil)
	if err != nil {
		return nil, err
	}

	response := []*BillerItem{}

	_, err = b.client.Perform(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Validate validates bill details for a customer.
// https://docs.swervpay.co/api-reference/bills/validate
func (b BillIntImpl) Validate(ctx context.Context, body *ValidateBillBody) error {
	req, err := b.client.NewRequest(ctx, http.MethodPost, "bills/validate", body)
	if err != nil {
		return err
	}

	_, err = b.client.Perform(req, nil)
	if err != nil {
		return err
	}

	return nil
}
