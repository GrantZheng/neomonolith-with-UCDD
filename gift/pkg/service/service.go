package service

import (
	"context"
)

// GiftService describes the service.
type GiftService interface {
	// Add your methods here
	Give(ctx context.Context, from string, tu string, gift_id string, gift_num string) (rs string, err error)
}

type basicGiftService struct{}

func (b *basicGiftService) Give(ctx context.Context, from string, tu string, gift_id string, gift_num string) (rs string, err error) {
	// TODO implement the business logic of Give
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
