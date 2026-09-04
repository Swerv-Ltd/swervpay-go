package swervpay

import (
	"context"
	"net/http"
)

// AuthInt defines authentication operations.
type AuthInt interface {
	Authenticate(ctx context.Context) (*AuthResponse, error)
}

// AuthIntImpl implements AuthInt.
type AuthIntImpl struct {
	client *SwervpayClient
}

var _ AuthInt = &AuthIntImpl{}

// Authenticate obtains an access token using the business credentials.
func (a AuthIntImpl) Authenticate(ctx context.Context) (*AuthResponse, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPost, "auth", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(a.client.Config.BusinessID, a.client.Config.SecretKey)
	response := new(AuthResponse)
	if _, err = a.client.Perform(req, response); err != nil {
		return nil, err
	}

	a.client.AccessToken = response.AccessToken
	return response, nil
}
