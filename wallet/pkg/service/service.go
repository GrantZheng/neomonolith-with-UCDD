package service

import (
	"context"
)

// WalletService describes the service.
type WalletService interface {
	// Add your methods here
	Deduct(ctx context.Context, uid string, money int) (rs string, err error)
}

type basicWalletService struct{}

func (b *basicWalletService) Deduct(ctx context.Context, uid string, money int) (rs string, err error) {
	// TODO implement the business logic of Deduct
	return rs, err
}

// NewBasicWalletService returns a naive, stateless implementation of WalletService.
func NewBasicWalletService() WalletService {
	return &basicWalletService{}
}

// New returns a WalletService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletService {
	var svc WalletService = NewBasicWalletService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
