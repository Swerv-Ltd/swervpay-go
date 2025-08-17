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
	CustomerID   string  `json:"customer_id"`   // The ID of the customer.
	Currency     string  `json:"currency"`      // The currency of the collection.
	MerchantName string  `json:"merchant_name"` // The name of the merchant.
	Amount       float64 `json:"amount"`        // The amount of the collection.
	Type         string  `json:"type"`          // The type of the collection.
}

// CollectionInt is an interface that defines the operations that can be performed on collections.
type CollectionInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Wallet, error)                               // Gets a list of wallets.
	Get(ctx context.Context, id string) (*Wallet, error)                                                 // Gets a specific wallet.
	Create(ctx context.Context, body *CreateCollectionBody) (*Wallet, error)                             // Creates a new wallet.
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
