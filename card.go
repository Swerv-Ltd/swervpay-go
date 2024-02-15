package main

import "context"

type Card struct{}

type CreateCardBody struct {
}

type FundOrWithdrawCardBody struct {
	Amount float64 `json:"amount"`
}

type CardInt interface {
	Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error)
	Get(ctx context.Context, id string) (*Card, error)
	Create(ctx context.Context, body CreateCardBody) (*DefaultResponse, error)
	Fund(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error)
	Withdraw(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error)
	Terminate(ctx context.Context, id string) (*DefaultResponse, error)
	Freeze(ctx context.Context, id string) (*DefaultResponse, error)
}

type CardIntImpl struct {
	client *SwervpayClient
}

var _ CardInt = &CardIntImpl{}

func (c CardIntImpl) Gets(ctx context.Context, query *PageAndLimitQuery) (*[]Card, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Get(ctx context.Context, id string) (*Card, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Create(ctx context.Context, body CreateCardBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Fund(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Withdraw(ctx context.Context, id string, body FundOrWithdrawCardBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Terminate(ctx context.Context, id string) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardIntImpl) Freeze(ctx context.Context, id string) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}
