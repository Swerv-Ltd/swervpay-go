package swervpay

import (
	"context"
	"net/http"
)

// BVNRequest represents a Bank Verification Number lookup request.
type BVNRequest struct {
	Number string `json:"number"`
}

// BvnRequest is kept as a MixedCaps spelling alias for callers that prefer the
// operation name from the API documentation.
type BvnRequest = BVNRequest

// BVNInput is an alias for BVNRequest.
type BVNInput = BVNRequest

// BvnInput is an alias for BVNRequest.
type BvnInput = BVNRequest

// BVNResponse represents the result of a Bank Verification Number lookup.
type BVNResponse struct {
	BVN                string `json:"bvn"`
	DateOfBirth        string `json:"date_of_birth"`
	Email              string `json:"email"`
	EnrollmentBank     string `json:"enrollment_bank"`
	EnrollmentBranch   string `json:"enrollment_branch"`
	FirstName          string `json:"first_name"`
	Gender             string `json:"gender"`
	LastName           string `json:"last_name"`
	LGAOfOrigin        string `json:"lga_of_Origin"`
	LGAOfResidence     string `json:"lga_of_residence"`
	MaritalStatus      string `json:"marital_status"`
	MiddleName         string `json:"middle_name"`
	Nationality        string `json:"nationality"`
	PhoneNumber        string `json:"phone_number"`
	Reference          string `json:"reference"`
	RegistrationDate   string `json:"registration_date"`
	ResidentialAddress string `json:"residential_address"`
	StateOfOrigin      string `json:"state_of_origin"`
	StateOfResidence   string `json:"state_of_residence"`
	Title              string `json:"title"`
}

// BvnResponse is an alias for BVNResponse.
type BvnResponse = BVNResponse

// IdentityInt defines identity verification operations.
type IdentityInt interface {
	BVN(ctx context.Context, body *BVNRequest) (*BVNResponse, error)
	Bvn(ctx context.Context, body *BvnRequest) (*BvnResponse, error)
}

// IdentityIntImpl implements IdentityInt.
type IdentityIntImpl struct {
	client *SwervpayClient
}

var _ IdentityInt = &IdentityIntImpl{}

// BVN verifies a Bank Verification Number.
func (i IdentityIntImpl) BVN(ctx context.Context, body *BVNRequest) (*BVNResponse, error) {
	req, err := i.client.NewRequest(ctx, http.MethodPost, "identity/bvn", body)
	if err != nil {
		return nil, err
	}

	response := new(BVNResponse)
	if _, err = i.client.Perform(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Bvn is an alias for BVN using the API operation's casing.
func (i IdentityIntImpl) Bvn(ctx context.Context, body *BvnRequest) (*BvnResponse, error) {
	return i.BVN(ctx, body)
}
