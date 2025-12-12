package swervpay

import (
	"context"
	"net/http"
)

// FxBody represents the body of a foreign exchange request.
type FxBody struct {
	Amount float64 `json:"amount"` // Amount is the amount to be converted.
	From   string  `json:"from"`   // From is the currency to convert from.
	To     string  `json:"to"`     // To is the currency to convert to.
}

// FxRateResponse represents the response from a foreign exchange rate request.
type FxRateResponse struct {
	Rate float64  `json:"rate"` // Rate is the conversion rate.
	From FromOrTo `json:"from"` // From represents the original currency and amount.
	To   FromOrTo `json:"to"`   // To represents the converted currency and amount.
}

// FromOrTo represents a currency and amount in a foreign exchange operation.
type FromOrTo struct {
	Amount   float64 `json:"amount"`   // Amount is the amount in the currency.
	Currency string  `json:"currency"` // Currency is the currency code.
}

// FxInt is an interface for foreign exchange operations.
type FxInt interface {
	// Rate gets the conversion rate for a foreign exchange operation.
	Rate(ctx context.Context, body FxBody) (*FxRateResponse, error)
	// Exchange performs a foreign exchange operation.
	Exchange(ctx context.Context, body FxBody) (*Transaction, error)
}

// FxIntImpl is an implementation of the FxInt interface.
type FxIntImpl struct {
	client *SwervpayClient // client is the Swervpay client.
}

// Verify that FxIntImpl implements FxInt.
var _ FxInt = &FxIntImpl{}

// Rate gets the conversion rate for a foreign exchange operation.
// https://docs.swervpay.co/api-reference/fx/get
func (f FxIntImpl) Rate(ctx context.Context, body FxBody) (*FxRateResponse, error) {
	req, err := f.client.NewRequest(ctx, http.MethodPost, "fx/rate", body)
	if err != nil {
		return nil, err
	}

	response := new(FxRateResponse)

	_, err = f.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Exchange performs a foreign exchange operation.
// https://docs.swervpay.co/api-reference/fx/create
func (f FxIntImpl) Exchange(ctx context.Context, body FxBody) (*Transaction, error) {
	req, err := f.client.NewRequest(ctx, http.MethodPost, "fx/exchange", body)
	if err != nil {
		return nil, err
	}

	response := new(Transaction)

	_, err = f.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
