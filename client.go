package swervpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	version     = "0.0.13"
	userAgent   = "Swervpay/Go-Sdk " + version
	contentType = "application/json"
)

// SwervpayClientOption represents the options for SwervpayClient.
type SwervpayClientOption struct {
	BusinessID string
	SecretKey  string
	Sandbox    bool
	Timeout    int
	Version    string
	BaseURL    string
}

type AuthResponse struct {
	AccessToken string      `json:"access_token"`
	Token       TokenDetail `json:"token"`
}

type TokenDetail struct {
	Type      string `json:"type"`
	ExpiresAt int64  `json:"expires_at"`
	IssuedAt  int64  `json:"issued_at"`
}

// SwervpayClient represents a client for interacting with Swervpay API Client.
type SwervpayClient struct {
	client      *http.Client
	Config      *SwervpayClientOption
	AccessToken string

	BaseURL *url.URL

	headers map[string]string

	Customer    CustomerInt
	Card        CardInt
	Business    BusinessInt
	Fx          FxInt
	Payout      PayoutInt
	Wallet      WalletInt
	Webhook     WebhookInt
	Transaction TransactionInt
	Other       OtherInt
	Collection  CollectionInt
	Bill        BillInt
}

// NewSwervpayClient creates a new SwervpayClient with the given options.
func NewSwervpayClient(config *SwervpayClientOption) *SwervpayClient {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.swervpay.co/api/v1/"

		if config.Sandbox {
			config.BaseURL = "https://sandbox.swervpay.co/api/v1/"
		}
	}
	baseURL, _ := url.Parse(config.BaseURL)

	s := &SwervpayClient{client: http.DefaultClient, Config: config, BaseURL: baseURL}

	s.Fx = &FxIntImpl{client: s}
	s.Business = &BusinessIntImpl{client: s}
	s.Transaction = &TransactionIntImpl{client: s}
	s.Payout = &PayoutIntImpl{client: s}
	s.Webhook = &WebhookIntImpl{client: s}
	s.Wallet = &WalletIntImpl{client: s}
	s.Other = &OtherIntImpl{client: s}
	s.Card = &CardIntImpl{client: s}
	s.Customer = &CustomerIntImpl{client: s}
	s.Collection = &CollectionIntImpl{client: s}
	s.Bill = &BillIntImpl{client: s}

	return s
}

// NewRequest creates a new request to the Swervpay API.
func (c *SwervpayClient) NewRequest(ctx context.Context, method, path string, params interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	if params != nil {
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(params)
		if err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(buf)
		req.Header.Set("Content-Type", contentType)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Accept", contentType)
	req.Header.Set("User-Agent", userAgent)

	if c.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	}

	return req, nil
}

// Perform sends the request to the Resend API
func (c *SwervpayClient) Perform(req *http.Request, ret interface{}) (*http.Response, error) {
	// Store the request body for potential retry after reauthentication
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body.Close()
		// Restore the body for the first attempt
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check if the status code is unauthorized
	if resp.StatusCode == http.StatusUnauthorized {
		// Make a request to the auth endpoint
		authReq, authReqErr := c.NewRequest(context.Background(), http.MethodPost, "auth", nil)
		if authReqErr != nil {
			return nil, authReqErr
		}

		authReq.SetBasicAuth(c.Config.BusinessID, c.Config.SecretKey)

		authResponse := new(AuthResponse)

		_, authReqErr = c.Perform(authReq, authResponse)
		if authReqErr != nil {
			return nil, authReqErr
		}

		// Update the access token
		c.AccessToken = authResponse.AccessToken

		// Update the original request's Authorization header
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)

		// Restore the request body from stored bytes for retry
		if bodyBytes != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Retry the request
		return c.Perform(req, ret)

	} else {

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)

		// Handle possible errors.
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			return nil, handleError(resp)
		}

		if resp.StatusCode != http.StatusNoContent && ret != nil {
			if w, ok := ret.(io.Writer); ok {
				_, err = io.Copy(w, resp.Body)
				if err != nil {
					return nil, err
				}
			} else {
				if resp.Body != nil {
					err = json.NewDecoder(resp.Body).Decode(ret)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		return resp, err
	}
}

func handleError(resp *http.Response) error {
	switch resp.StatusCode {

	// Handles errors most likely caused by the client
	case http.StatusUnprocessableEntity, http.StatusBadRequest:
		r := &InvalidRequestError{}
		err := json.NewDecoder(resp.Body).Decode(r)
		if err != nil {
			return err
		}
		return errors.New("[ERROR]: " + r.Message)
	case http.StatusNotFound:
		return errors.New("[ERROR]: Not Found")
	default:
		// Tries to parse `message` attr from error
		r := &DefaultResponse{}
		err := json.NewDecoder(resp.Body).Decode(r)
		if err != nil {
			return err
		}
		if r.Message != "" {
			return errors.New("[ERROR]: " + r.Message)
		}
		return errors.New("[ERROR]: Unknown Error")
	}
}

// InvalidRequestError represents an error caused by the client.
type InvalidRequestError struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Name       string `json:"name,omitempty"`
	Message    string `json:"message"`
}

// DefaultResponse represents the default response from the Swervpay API.
type DefaultResponse struct {
	Message string `json:"message"`
}
