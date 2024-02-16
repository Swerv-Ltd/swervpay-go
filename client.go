package main

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
	version     = "0.0.1"
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
}

// NewSwervpayClient creates a new SwervpayClient with the given options.
func NewSwervpayClient(config *SwervpayClientOption) *SwervpayClient {
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

	return s
}

func (c *SwervpayClient) NewRequest(ctx context.Context, method, path string, params interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
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
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	return req, nil
}

// Perform sends the request to the Resend API
func (c *SwervpayClient) Perform(req *http.Request, ret interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

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

type InvalidRequestError struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
}

type DefaultResponse struct {
	Message string `json:"message"`
}
