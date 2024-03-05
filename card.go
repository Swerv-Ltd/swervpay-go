package swervpay

import (
	"context"
	"net/http"
)

type Card struct {
	AddressCity       string `json:"address_city"`
	AddressCountry    string `json:"address_country"`
	AddressPostalCode string `json:"address_postal_code"`
	AddressState      string `json:"address_state"`
	AddressStreet     string `json:"address_street"`
	Balance           int64  `json:"balance"`
	CardNumber        string `json:"card_number"`
	CreatedAt         string `json:"created_at"`
	Currency          string `json:"currency"`
	Cvv               string `json:"cvv"`
	Expiry            string `json:"expiry"`
	Freeze            bool   `json:"freeze"`
	ID                string `json:"id"`
	Issuer            string `json:"issuer"`
	MaskedPan         string `json:"masked_pan"`
	NameOnCard        string `json:"name_on_card"`
	Status            string `json:"status"`
	TotalFunded       int64  `json:"total_funded"`
	Type              string `json:"type"`
	UpdatedAt         string `json:"updated_at"`
}

type CardTransactionHistory struct {
	Amount             int64  `json:"amount"`
	Category           string `json:"category"`
	Charges            int64  `json:"charges"`
	CreatedAt          string `json:"created_at"`
	Currency           string `json:"currency"`
	ID                 string `json:"id"`
	MerchantCity       string `json:"merchant_city"`
	MerchantCountry    string `json:"merchant_country"`
	MerchantMcc        string `json:"merchant_mcc"`
	MerchantMid        string `json:"merchant_mid"`
	MerchantName       string `json:"merchant_name"`
	MerchantPostalCode string `json:"merchant_postal_code"`
	MerchantState      string `json:"merchant_state"`
	Reference          string `json:"reference"`
	Report             bool   `json:"report"`
	ReportMessage      string `json:"report_message"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	UpdatedAt          string `json:"updated_at"`
}

type CreateCardBody struct {
	Amount     float64 `json:"amount"`
	CustomerId string  `json:"customer_id"`
	Provider   string  `json:"provider"`
	NameOnCard string  `json:"name_on_card"`
	Currency   string  `json:"currency"`
	Type       string  `json:"type"`
}

type CardCreationResponse struct {
	CardID  string `json:"card_id"`
	Message string `json:"message"`
}

type FundOrWithdrawCardBody struct {
	Amount float64 `json:"amount"`
}

type CardInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error)
	Get(ctx context.Context, id string) (*Card, error)
	Create(ctx context.Context, body *CreateCardBody) (*CardCreationResponse, error)
	Fund(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error)
	Withdraw(ctx context.Context, id string, body *FundOrWithdrawCardBody) (*DefaultResponse, error)
	Terminate(ctx context.Context, id string) (*DefaultResponse, error)
	Freeze(ctx context.Context, id string) (*DefaultResponse, error)
	Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CardTransactionHistory, error)
	Transaction(ctx context.Context, id string, transactionId string) (*CardTransactionHistory, error)
}

type CardIntImpl struct {
	client *SwervpayClient
}

var _ CardInt = &CardIntImpl{}

func (c CardIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error) {
	path := GenerateURLPath("cards", query)

	// Prepare request
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

func (c CardIntImpl) Transactions(ctx context.Context, id string, query *PageAndLimitQuery) (*[]CardTransactionHistory, error) {
	path := GenerateURLPath("cards/"+id+"/transactions", query)

	// Prepare request
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

func (c CardIntImpl) Transaction(ctx context.Context, id string, transactionId string) (*CardTransactionHistory, error) {

	// Prepare request
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
