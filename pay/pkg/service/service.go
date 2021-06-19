package service

import (
	"context"
)

// PayService describes the service.
type PayService interface {
	// Add your methods here
	Pay(ctx context.Context, account float32, user_id string) (rs string, err error)
}

type basicPayService struct{}

func (b *basicPayService) Pay(ctx context.Context, account float32, user_id string) (rs string, err error) {
	// TODO implement the business logic of Pay
	return rs, err
}

// NewBasicPayService returns a naive, stateless implementation of PayService.
func NewBasicPayService() PayService {
	return &basicPayService{}
}

// New returns a PayService with all of the expected middleware wired in.
func New(middleware []Middleware) PayService {
	var svc PayService = NewBasicPayService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
