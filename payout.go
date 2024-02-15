package main

import "context"

type CreatePayoutBody struct {
}

type PayoutInt interface {
	Get(ctx context.Context, id string) (*Transaction, error)
	Create(ctx context.Context, body CreatePayoutBody) (*DefaultResponse, error)
}

type PayoutIntImpl struct {
	client *SwervpayClient
}

var _ PayoutInt = &PayoutIntImpl{}

func (p PayoutIntImpl) Get(ctx context.Context, id string) (*Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (p PayoutIntImpl) Create(ctx context.Context, body CreatePayoutBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}
