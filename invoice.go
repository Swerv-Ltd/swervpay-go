package swervpay

import (
	"context"
	"net/http"
)

// Invoice represents an invoice returned by the API.
type Invoice struct {
	Attachments         string        `json:"attachments"`
	Categories          string        `json:"categories"`
	CreatedAt           string        `json:"created_at"`
	Currency            string        `json:"currency"`
	Description         string        `json:"description"`
	DueDate             string        `json:"due_date"`
	ID                  string        `json:"id"`
	InvoiceLink         string        `json:"invoice_link"`
	InvoiceNumber       string        `json:"invoice_number"`
	IsRecurring         bool          `json:"is_recurring"`
	IssuedDate          string        `json:"issued_date"`
	Items               []InvoiceItem `json:"items"`
	Memo                string        `json:"memo"`
	RecurringEndDate    string        `json:"recurring_end_date"`
	RecurringFrequency  string        `json:"recurring_frequency"`
	Status              string        `json:"status"`
	Tags                string        `json:"tags"`
	Title               string        `json:"title"`
	TotalAmount         float64       `json:"total_amount"`
	TotalDiscountAmount float64       `json:"total_discount_amount"`
	TotalTaxFee         float64       `json:"total_tax_fee"`
	UpdatedAt           string        `json:"updated_at"`
}

// InvoiceItem represents an invoice line item returned by the API.
type InvoiceItem struct {
	CreatedAt      string  `json:"created_at"`
	Description    string  `json:"description"`
	DiscountAmount float64 `json:"discount_amount"`
	ID             string  `json:"id"`
	Quantity       int     `json:"quantity"`
	Tax            string  `json:"tax"`
	TaxFee         float64 `json:"tax_fee"`
	TotalPrice     float64 `json:"total_price"`
	UnitPrice      float64 `json:"unit_price"`
	UpdatedAt      string  `json:"updated_at"`
}

// InvoiceItemBody represents an invoice line item request.
type InvoiceItemBody struct {
	Description string  `json:"description"`
	Discount    float64 `json:"discount"`
	ID          string  `json:"id"`
	Quantity    int     `json:"quantity"`
	Tax         string  `json:"tax"`
	TaxFee      float64 `json:"tax_fee"`
	UnitPrice   float64 `json:"unit_price"`
}

// CreateInvoiceItemBody is an alias for InvoiceItemBody.
type CreateInvoiceItemBody = InvoiceItemBody

// CreateInvoiceItemInput is an alias for InvoiceItemBody.
type CreateInvoiceItemInput = InvoiceItemBody

// CreateInvoiceBody represents an invoice creation request.
type CreateInvoiceBody struct {
	Categories               []string          `json:"categories"`
	CC                       []string          `json:"cc"`
	Currency                 string            `json:"currency"`
	CustomerBillingAddressID string            `json:"customer_billing_address_id"`
	CustomerID               string            `json:"customer_id"`
	Description              string            `json:"description"`
	DueDate                  string            `json:"due_date"`
	IsRecurring              bool              `json:"is_recurring"`
	IssueDate                string            `json:"issue_date"`
	Items                    []InvoiceItemBody `json:"items"`
	Memo                     string            `json:"memo"`
	Recurring                string            `json:"recurring"`
	Status                   string            `json:"status"`
	Tags                     []string          `json:"tags"`
	Title                    string            `json:"title"`
}

// UpdateInvoiceBody represents an invoice update request.
type UpdateInvoiceBody = CreateInvoiceBody

// InvoiceInt defines invoice operations.
type InvoiceInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Invoice, error)
	Get(ctx context.Context, id string) (*Invoice, error)
	Create(ctx context.Context, body *CreateInvoiceBody) (*Invoice, error)
	Update(ctx context.Context, id string, body *UpdateInvoiceBody) (*Invoice, error)
}

// InvoiceIntImpl implements InvoiceInt.
type InvoiceIntImpl struct {
	client *SwervpayClient
}

var _ InvoiceInt = &InvoiceIntImpl{}

// Gets lists invoices.
func (i InvoiceIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Invoice, error) {
	path := GenerateURLPath("invoices", query)
	req, err := i.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*Invoice{}
	if _, err = i.client.Perform(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get retrieves an invoice by ID.
func (i InvoiceIntImpl) Get(ctx context.Context, id string) (*Invoice, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, "invoices/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Invoice)
	if _, err = i.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create creates an invoice.
func (i InvoiceIntImpl) Create(ctx context.Context, body *CreateInvoiceBody) (*Invoice, error) {
	req, err := i.client.NewRequest(ctx, http.MethodPost, "invoices", body)
	if err != nil {
		return nil, err
	}

	response := new(Invoice)
	if _, err = i.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update updates an invoice by ID.
func (i InvoiceIntImpl) Update(ctx context.Context, id string, body *UpdateInvoiceBody) (*Invoice, error) {
	req, err := i.client.NewRequest(ctx, http.MethodPost, "invoices/"+id, body)
	if err != nil {
		return nil, err
	}

	response := new(Invoice)
	if _, err = i.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}
