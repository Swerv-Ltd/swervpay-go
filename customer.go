package swervpay

import (
	"context"
	"net/http"
)

// Customer represents a customer in the Swervpay system.
type Customer struct {
	Country       string `json:"country"`        // The country of the customer.
	CreatedAt     string `json:"created_at"`     // The creation date of the customer.
	Email         string `json:"email"`          // The email of the customer.
	FirstName     string `json:"first_name"`     // The first name of the customer.
	ID            string `json:"id"`             // The ID of the customer.
	IsBlacklisted bool   `json:"is_blacklisted"` // Whether the customer is blacklisted.
	LastName      string `json:"last_name"`      // The last name of the customer.
	MiddleName    string `json:"middle_name"`    // The middle name of the customer.
	PhoneNumber   string `json:"phone_number"`   // The phone number of the customer.
	Status        string `json:"status"`         // The status of the customer.
	UpdatedAt     string `json:"updated_at"`     // The last update date of the customer.
}

// CreateCustomerBody represents the body of a request to create a new customer.
type CreateCustomerBody struct {
	Country    string `json:"country"`    // The country of the new customer.
	Email      string `json:"email"`      // The email of the new customer.
	Firstname  string `json:"firstname"`  // The first name of the new customer.
	Lastname   string `json:"lastname"`   // The last name of the new customer.
	Middlename string `json:"middlename"` // The middle name of the new customer.
}

// UpdateustomerBody represents the body of a request to update a customer.
type UpdateustomerBody struct {
	Email       string `json:"email"`        // The new email of the customer.
	PhoneNumber string `json:"phone_number"` // The new phone number of the customer.
}

// CustomerKycBody represents the body of a request to update a customer's KYC information.
type CustomerKycBody struct {
	Tier  string        `json:"tier"`        // The tier of the KYC information.
	Tier1 Tier1KycInput `json:"information"` // The tier 1 KYC information.
	Tier2 Tier2KycInput `json:"document"`    // The tier 2 KYC information.
}

// Tier1KycInput represents the tier 1 KYC information of a customer.
type Tier1KycInput struct {
	Bvn         string `json:"bvn"`          // The BVN of the customer.
	State       string `json:"state"`        // The state of the customer.
	City        string `json:"city"`         // The city of the customer.
	Country     string `json:"country"`      // The country of the customer.
	Address     string `json:"address"`      // The address of the customer.
	PostalCode  string `json:"postal_code"`  // The postal code of the customer.
	PhoneNumber string `json:"phone_number"` // The phone number of the customer.
}

// Tier2KycInput represents the tier 2 KYC information of a customer.
type Tier2KycInput struct {
	DocumentType   string `json:"document_type"`   // The type of the document.
	Document       string `json:"document"`        // The document.
	Passport       string `json:"passport"`        // The passport of the customer.
	DocumentNumber string `json:"document_number"` // The document number.
}

// CustomerInt is an interface that defines the methods for interacting with customers in the Swervpay system.
type CustomerInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Customer, error)             // Gets a list of customers.
	Get(ctx context.Context, id string) (*Customer, error)                               // Gets a specific customer.
	Create(ctx context.Context, body *CreateCustomerBody) (*Customer, error)             // Creates a new customer.
	Update(ctx context.Context, id string, body *UpdateustomerBody) (*Customer, error)   // Updates a specific customer.
	Kyc(ctx context.Context, id string, body *CustomerKycBody) (*DefaultResponse, error) // Updates the KYC information of a specific customer.
	Blacklist(ctx context.Context, id string) (*DefaultResponse, error)                  // Blacklists a specific customer.
}

// CustomerIntImpl is an implementation of the CustomerInt interface.
type CustomerIntImpl struct {
	client *SwervpayClient // The Swervpay client.
}

// Verify that CustomerIntImpl implements the CustomerInt interface.
var _ CustomerInt = &CustomerIntImpl{}

// Gets retrieves a list of customers.
// https://docs.swervpay.co/api-reference/customers/get-all-customers
func (c CustomerIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) ([]*Customer, error) {
	path := GenerateURLPath("customers", query)

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := []*Customer{}

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get retrieves a specific customer.
// https://docs.swervpay.co/api-reference/customers/get
func (c CustomerIntImpl) Get(ctx context.Context, id string) (*Customer, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "customers/"+id, nil)
	if err != nil {
		return nil, err
	}

	response := new(Customer)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Create creates a new customer.
// https://docs.swervpay.co/api-reference/customers/create
func (c CustomerIntImpl) Create(ctx context.Context, body *CreateCustomerBody) (*Customer, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "customers", body)
	if err != nil {
		return nil, err
	}

	response := new(Customer)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Update updates a specific customer.
// https://docs.swervpay.co/api-reference/customers/update
func (c CustomerIntImpl) Update(ctx context.Context, id string, body *UpdateustomerBody) (*Customer, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "customers/"+id+"/update", body)
	if err != nil {
		return nil, err
	}

	response := new(Customer)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Kyc updates the KYC information of a specific customer.
// https://docs.swervpay.co/api-reference/customers/kyc
func (c CustomerIntImpl) Kyc(ctx context.Context, id string, body *CustomerKycBody) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "customers/"+id+"/kyc", body)
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

// Blacklist blacklists a specific customer.
// https://docs.swervpay.co/api-reference/customers/blacklist
func (c CustomerIntImpl) Blacklist(ctx context.Context, id string) (*DefaultResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "customers/"+id+"/blacklist", nil)
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
