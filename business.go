// Package swervpay provides a set of APIs to interact with the Swervpay service.
package swervpay

import (
	"context"
	"net/http"
)

// Business represents a business entity in the Swervpay system.
// It includes various properties like address, name, country, etc.
type Business struct {
	Address   string `json:"address"`    // Address of the business
	Name      string `json:"name"`       // Name of the business
	Country   string `json:"country"`    // Country where the business is located
	CreatedAt string `json:"created_at"` // Time when the business was created
	Email     string `json:"email"`      // Email of the business
	ID        string `json:"id"`         // Unique identifier of the business
	Logo      string `json:"logo"`       // Logo of the business
	Slug      string `json:"slug"`       // Slug of the business
	Type      string `json:"type"`       // Type of the business
	UpdatedAt string `json:"updated_at"` // Time when the business was last updated
}

// BusinessInt is an interface that defines the methods a Business must have.
type BusinessInt interface {
	// Get retrieves the Business information from the Swervpay system.
	Get(ctx context.Context) (*Business, error)
}

// BusinessIntImpl is a concrete implementation of the BusinessInt interface.
type BusinessIntImpl struct {
	client *SwervpayClient // Client used to make requests to the Swervpay system
}

// Ensure BusinessIntImpl implements BusinessInt interface
var _ BusinessInt = &BusinessIntImpl{}

// Get is a method on BusinessIntImpl that retrieves the Business information from the Swervpay system.
// https://docs.swervpay.co/api-reference/business/get
func (b *BusinessIntImpl) Get(ctx context.Context) (*Business, error) {

	// Create a new request to get the business information
	req, err := b.client.NewRequest(ctx, http.MethodGet, "business", nil)
	if err != nil {
		return nil, err
	}

	// Response will hold the business information
	response := new(Business)

	// Perform the request and populate the response
	_, err = b.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	// Return the business information
	return response, nil
}
