package main

import "context"

type FxBody struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type FxInt interface {
	Rate(ctx context.Context, body FxBody) (*DefaultResponse, error)
	Exchange(ctx context.Context, body FxBody) (*DefaultResponse, error)
}

type FxIntImpl struct {
	client *SwervpayClient
}

var _ FxInt = &FxIntImpl{}

func (f FxIntImpl) Rate(ctx context.Context, body FxBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FxIntImpl) Exchange(ctx context.Context, body FxBody) (*DefaultResponse, error) {
	//TODO implement me
	panic("implement me")
}
