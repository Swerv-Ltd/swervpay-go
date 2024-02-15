package main

import "context"

type Customer struct{}

type CreateCustomerBody struct {
}

type UpdateustomerBody struct {
}

type CustomerKycBody struct {
}

type CustomerInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Customer, error)
	Get(ctx context.Context, id string) (*Customer, error)
	Create(ctx context.Context, body CreateCustomerBody) (*Customer, error)
	Update(ctx context.Context, id string, body UpdateustomerBody) (*Customer, error)
	Kyc(ctx context.Context, id string, body CustomerKycBody) (*DefaultResponse, error)
	Blacklist(ctx context.Context, id string) (*DefaultResponse, error)
}

type CustomerIntImpl struct {
	client *SwervpayClient
}

var _ CustomerInt = &CustomerIntImpl{}

func (c CustomerIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerIntImpl) Get(ctx context.Context, id string) (*Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerIntImpl) Create(ctx context.Context, body CreateCustomerBody) (*Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerIntImpl) Update(ctx context.Context, id string, body UpdateustomerBody) (*Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerIntImpl) Kyc(ctx context.Context, id string, body CustomerKycBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerIntImpl) Blacklist(ctx context.Context, id string) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}
