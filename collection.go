package swervpay

import (
	"context"
	"net/http"
)

// CollectionHistory represents the history of a collection.
type CollectionHistory struct {
	Amount        float64 `json:"amount"`         // The amount of the collection.
	Charges       float64 `json:"charges"`        // The charges associated with the collection.
	CreatedAt     string  `json:"created_at"`     // The creation date of the collection.
	Currency      string  `json:"currency"`       // The currency of the collection.
	ID            string  `json:"id"`             // The ID of the collection.
	PaymentMethod string  `json:"payment_method"` // The payment method used for the collection.
	Reference     string  `json:"reference"`      // The reference of the collection.
	UpdatedAt     string  `json:"updated_at"`     // The last update date of the collection.
}

// CreateCollectionBody represents the body of a create collection request.
type CreateCollectionBody struct {
	CustomerID            string                     `json:"customer_id"`                      // The ID of the customer.
	Currency              string                     `json:"currency"`                         // The currency of the collection.
	MerchantName          string                     `json:"merchant_name"`                    // The name of the merchant.
	Amount                float64                    `json:"amount"`                           // The amount of the collection.
	Type                  string                     `json:"type"`                             // The type of the collection.
	Reference             string                     `json:"reference,omitempty"`              // Optional reference for idempotency.
	AdditionalInformation *AdditionalInformationBody `json:"additional_information,omitempty"` // Optional additional information.
}

// AdditionalInformationBody mirrors TypesAdditionalInformation from the OpenAPI spec.
type AdditionalInformationBody struct {
	AccountDesignation string                     `json:"account_designation,omitempty"`
	AccountType        string                     `json:"account_type,omitempty"`
	Address            *AdditionalInformationAddr `json:"address,omitempty"`
	BankStatement      string                     `json:"bank_statement,omitempty"`
	DateOfBirth        string                     `json:"date_of_birth,omitempty"`
	Document           *AdditionalInformationDoc  `json:"document,omitempty"`
	EmploymentStatus   string                     `json:"employment_status,omitempty"`
	IncomeBand         string                     `json:"income_band,omitempty"`
	Nin                string                     `json:"nin,omitempty"`
	SourceOfIncome     string                     `json:"source_of_income,omitempty"`
	TaxNumber          string                     `json:"tax_number,omitempty"`
	UtilityBill        string                     `json:"utility_bill,omitempty"`
}

// AdditionalInformationAddr mirrors TypesAddress.
type AdditionalInformationAddr struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	Street  string `json:"street,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}

// AdditionalInformationDoc mirrors TypesDocument.
type AdditionalInformationDoc struct {
	ExpiryDate string   `json:"expiry_date,omitempty"`
	IssueDate  string   `json:"issue_date,omitempty"`
	Number     string   `json:"number,omitempty"`
	Type       string   `json:"type,omitempty"`
	URLs       []string `json:"urls,omitempty"`
}

// CollectionInt is an interface that defines the operations that can be performed on collections.
type CollectionInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Wallet, error)                               // Gets a list of wallets.
	Get(ctx context.Context, id string) (*Wallet, error)                                                 // Gets a specific wallet.
	Create(ctx context.Context, body *CreateCollectionBody) (*Wallet, error)                             // Creates a new wallet.
	Credit(ctx context.Context, id string, body *CreditWalletBody) (*CreditWalletResponse, error)        // Credits a collection.
	Transactions(ctx context.Context, id string, query *PageAndLimitQuery) ([]*CollectionHistory, error) // Gets the transactions of a specific wallet.
}

// CollectionIntImpl is an implementation of the CollectionInt interface.
type CollectionIntImpl struct {
	client *SwervpayClient // The client used to interact with the Swervpay API.
}

// Verify that CollectionIntImpl implements CollectionInt.
var _ CollectionInt = &CollectionIntImpl{}

// Gets retrieves a list of wallets.
// https://docs.swervpay.co/api-reference/collections/get-all-collections
func (c CollectionIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Wallet, error) {
	path := GenerateURLPath("collections", query)

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*Wallet{}

	_, err = c.client.Perform(req, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get retrieves a specific wallet.
// https://docs.swervpay.co/api-reference/collections/get
func (c CollectionIntImpl) Get(ctx context.Context, id string) (*Wallet, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Create creates a new wallet.
// https://docs.swervpay.co/api-reference/collections/create
func (c CollectionIntImpl) Create(ctx context.Context, body *CreateCollectionBody) (*Wallet, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "collections", body)
	if err != nil {
		return nil, err
	}

	response := new(Wallet)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Credit credits a collection.
// https://docs.swervpay.co/api-reference/collections/credit
func (c CollectionIntImpl) Credit(ctx context.Context, id string, body *CreditWalletBody) (*CreditWalletResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "collections/"+id+"/credit", body)
	if err != nil {
		return nil, err
	}

	response := new(CreditWalletResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Transactions retrieves the transactions of a specific wallet.
// https://docs.swervpay.co/api-reference/collections/transaction
func (c CollectionIntImpl) Transactions(ctx context.Context, id string, query *PageAndLimitQuery) ([]*CollectionHistory, error) {
	path := GenerateURLPath("collections/"+id+"/transactions", query)

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*CollectionHistory{}

	_, err = c.client.Perform(req, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
