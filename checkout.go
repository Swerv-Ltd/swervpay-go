package swervpay

import (
	"context"
	"net/http"
)

// CreditCheckoutBody represents the payload for a sandbox checkout credit.
type CreditCheckoutBody = CreditWalletBody

// CreditCheckoutResponse represents the result of a sandbox checkout credit.
type CreditCheckoutResponse = CreatePayoutResponse

// CheckoutInt defines checkout operations.
type CheckoutInt interface {
	// Credit simulates a credit transaction for a checkout account.
	Credit(ctx context.Context, accountNo string, body *CreditCheckoutBody) (*CreditCheckoutResponse, error)
}

// CheckoutIntImpl implements CheckoutInt.
type CheckoutIntImpl struct {
	client *SwervpayClient
}

var _ CheckoutInt = &CheckoutIntImpl{}

// Credit simulates a checkout credit transaction. This endpoint is available in
// the sandbox environment only.
func (c CheckoutIntImpl) Credit(ctx context.Context, accountNo string, body *CreditCheckoutBody) (*CreditCheckoutResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "checkouts/"+accountNo+"/credit", body)
	if err != nil {
		return nil, err
	}

	response := new(CreditCheckoutResponse)
	if _, err = c.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}
