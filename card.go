package swervpay

import (
	"context"
	"net/http"
)

// Card represents a card with its details.
type Card struct {
	AddressCity       string `json:"address_city"`        // City of the card holder's address.
	AddressCountry    string `json:"address_country"`     // Country of the card holder's address.
	AddressPostalCode string `json:"address_postal_code"` // Postal code of the card holder's address.
	AddressState      string `json:"address_state"`       // State of the card holder's address.
	AddressStreet     string `json:"address_street"`      // Street of the card holder's address.
	Balance           int64  `json:"balance"`             // Balance on the card.
	CardNumber        string `json:"card_number"`         // Card number.
	CreatedAt         string `json:"created_at"`          // Creation date of the card.
	Currency          string `json:"currency"`            // Currency of the card.
	Cvv               string `json:"cvv"`                 // CVV of the card.
	Expiry            string `json:"expiry"`              // Expiry date of the card.
	Freeze            bool   `json:"freeze"`              // Freeze status of the card.
	ID                string `json:"id"`                  // ID of the card.
	Issuer            string `json:"issuer"`              // Issuer of the card.
	MaskedPan         string `json:"masked_pan"`          // Masked PAN of the card.
	NameOnCard        string `json:"name_on_card"`        // Name on the card.
	Status            string `json:"status"`              // Status of the card.
	TotalFunded       int64  `json:"total_funded"`        // Total funded amount on the card.
	Type              string `json:"type"`                // Type of the card.
	UpdatedAt         string `json:"updated_at"`          // Last update date of the card.
}

// CardTransactionHistory represents a card's transaction history.
type CardTransactionHistory struct {
	Amount             int64  `json:"amount"`               // Transaction amount.
	Category           string `json:"category"`             // Category of the transaction.
	Charges            int64  `json:"charges"`              // Charges of the transaction.
	CreatedAt          string `json:"created_at"`           // Creation date of the transaction.
	Currency           string `json:"currency"`             // Currency of the transaction.
	ID                 string `json:"id"`                   // ID of the transaction.
	MerchantCity       string `json:"merchant_city"`        // City of the merchant.
	MerchantCountry    string `json:"merchant_country"`     // Country of the merchant.
	MerchantMcc        string `json:"merchant_mcc"`         // MCC of the merchant.
	MerchantMid        string `json:"merchant_mid"`         // MID of the merchant.
	MerchantName       string `json:"merchant_name"`        // Name of the merchant.
	MerchantPostalCode string `json:"merchant_postal_code"` // Postal code of the merchant.
	MerchantState      string `json:"merchant_state"`       // State of the merchant.
	Reference          string `json:"reference"`            // Reference of the transaction.
	Report             bool   `json:"report"`               // Report status of the transaction.
	ReportMessage      string `json:"report_message"`       // Report message of the transaction.
	Status             string `json:"status"`               // Status of the transaction.
	Type               string `json:"type"`                 // Type of the transaction.
	UpdatedAt          string `json:"updated_at"`           // Last update date of the transaction.
}

// CreateCardBody represents the body of a card creation request.
type CreateCardBody struct {
	Amount     float64 `json:"amount"`       // Amount to be loaded on the card.
	CustomerId string  `json:"customer_id"`  // ID of the customer.
	Provider   string  `json:"provider"`     // Provider of the card.
	NameOnCard string  `json:"name_on_card"` // Name to be printed on the card.
	Currency   string  `json:"currency"`     // Currency of the card.
	Type       string  `json:"type"`         // Type of the card.
}

// CardCreationResponse represents the response of a card creation request.
type CardCreationResponse struct {
	CardID  string `json:"card_id"` // ID of the created card.
	Message string `json:"message"` // Message of the response.
}

// FundOrWithdrawCardBody represents the body of a fund or withdraw request.
type FundOrWithdrawCardBody struct {
	Amount float64 `json:"amount"` // Amount to be funded or withdrawn.
}

// CardInt is the interface for card operations.
type CardInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error)                                      // Gets multiple cards.
	Get(ctx context.Context, id string) (*Card, error)                                                        // Gets a single card.
	Create(ctx context.Context, body *CreateCardBody) (*CardCreationResponse, error)                          // Creates a card.
	Fund(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error)              // Funds a card.
	Withdraw(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error)          // Withdraws from a card.
	Terminate(ctx context.Context, id string) (*DefaultResponse, error)                                       // Terminates a card.
	Freeze(ctx context.Context, id string) (*DefaultResponse, error)                                          // Freezes a card.
	Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CardTransactionHistory, error) // Gets multiple transactions of a card.
	Transaction(ctx context.Context, id string, transactionId string) (*CardTransactionHistory, error)        // Gets a single transaction of a card.
}

// CardIntImpl is the implementation of the CardInt interface.
type CardIntImpl struct {
	client *SwervpayClient // Client to perform the requests.
}

// Verify that CardIntImpl implements CardInt.
var _ CardInt = &CardIntImpl{}

// Gets gets multiple cards.
// https://docs.swervpay.co/api-reference/cards/get-all-cards
func (c CardIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error) {
	path := GenerateURLPath("cards", query)

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Card)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets a single card.
// https://docs.swervpay.co/api-reference/cards/get
func (c CardIntImpl) Get(ctx context.Context, id string) (*Card, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "cards/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Card)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Create creates a card.
// https://docs.swervpay.co/api-reference/cards/create
func (c CardIntImpl) Create(ctx context.Context, body *CreateCardBody) (*CardCreationResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards", body)
	if err != nil {
		return nil, err
	}

	response := new(CardCreationResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Fund funds a card.
// https://docs.swervpay.co/api-reference/cards/fund-card
func (c CardIntImpl) Fund(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/fund", body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Withdraw withdraws from a card.
// https://docs.swervpay.co/api-reference/cards/withdraw-from-card
func (c CardIntImpl) Withdraw(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/withdraw", body)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Terminate terminates a card.
// https://docs.swervpay.co/api-reference/cards/terminate
func (c CardIntImpl) Terminate(ctx context.Context, id string) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/terminate", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Freeze freezes a card.
// https://docs.swervpay.co/api-reference/cards/freeze
func (c CardIntImpl) Freeze(ctx context.Context, id string) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "cards/"+id+"/freeze", nil)
	if err != nil {
		return nil, err
	}

	response := new(DefaultResponse)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Transactions gets multiple transactions of a card.
// https://docs.swervpay.co/api-reference/cards/transactions
func (c CardIntImpl) Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CardTransactionHistory, error) {
	path := GenerateURLPath("cards/"+id+"/transactions", query)

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]CardTransactionHistory)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Transaction get transactions of a card.
// https://docs.swervpay.co/api-reference/cards/get-transaction
func (c CardIntImpl) Transaction(ctx context.Context, id string, transactionId string) (*CardTransactionHistory, error) {

	req, err := c.client.NewRequest(ctx, http.MethodGet, "cards/"+id+"/transactions/"+transactionId, nil)
	if err != nil {
		return nil, err
	}

	response := new(CardTransactionHistory)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
