package service

import (
	"context"
)

// GiftService describes the service.
type GiftService interface {
	// Add your methods here
	List(ctx context.Context, s string) (rs string, err error)
	Send(ctx context.Context, from string, to string, gift_id string) (rs string, err error)
}

type basicGiftService struct{}

func (b *basicGiftService) List(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of List
	return rs, err
}
func (b *basicGiftService) Send(ctx context.Context, from string, to string, gift_id string) (rs string, err error) {
	// TODO implement the business logic of Send
	return rs, err
}

// NewBasicGiftService returns a naive, stateless implementation of GiftService.
func NewBasicGiftService() GiftService {
	return &basicGiftService{}
}

// New returns a GiftService with all of the expected middleware wired in.
func New(middleware []Middleware) GiftService {
	var svc GiftService = NewBasicGiftService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
