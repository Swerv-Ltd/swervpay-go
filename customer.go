package swervpay

import (
	"context"
	"net/http"
)

type Customer struct {
	BusinessID    string `json:"business_id"`
	Country       string `json:"country"`
	CreatedAt     string `json:"created_at"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	ID            string `json:"id"`
	IsBlacklisted bool   `json:"is_blacklisted"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	PhoneNumber   string `json:"phone_number"`
	Status        string `json:"status"`
	UpdatedAt     string `json:"updated_at"`
}

type CreateCustomerBody struct {
	Country    string `json:"country"`
	Email      string `json:"email"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type UpdateustomerBody struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerKycBody struct {
	Tier  string        `json:"tier"`
	Tier1 Tier1KycInput `json:"information"`
	Tier2 Tier2KycInput `json:"document"`
}

type Tier1KycInput struct {
	Bvn        string `json:"bvn"`
	State      string `json:"state"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}

type Tier2KycInput struct {
	DocumentType   string `json:"document_type"`
	Document       string `json:"document"`
	Passport       string `json:"passport"`
	DocumentNumber string `json:"document_number"`
}

type CustomerInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Customer, error)
	Get(ctx context.Context, id string) (*Customer, error)
	Create(ctx context.Context, body *CreateCustomerBody) (*Customer, error)
	Update(ctx context.Context, id string, body *UpdateustomerBody) (*Customer, error)
	Kyc(ctx context.Context, id string, body *CustomerKycBody) (*DefaultResponse, error)
	Blacklist(ctx context.Context, id string) (*DefaultResponse, error)
}

type CustomerIntImpl struct {
	client *SwervpayClient
}

var _ CustomerInt = &CustomerIntImpl{}

func (c CustomerIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Customer, error) {
	path := GenerateURLPath("customers", query)

	// Prepare request
	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]Customer)

	_, err = c.client.Perform(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

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
