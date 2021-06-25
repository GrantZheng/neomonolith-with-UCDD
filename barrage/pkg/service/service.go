package service

import (
	"context"
)

// BarrageService describes the service.
type BarrageService interface {
	// Add your methods here
	Send(ctx context.Context, from string, to string, msg string) (rs string, err error)
}

type basicBarrageService struct{}

func (b *basicBarrageService) Send(ctx context.Context, from string, to string, msg string) (rs string, err error) {
	// TODO implement the business logic of Send
	return rs, err
}

// NewBasicBarrageService returns a naive, stateless implementation of BarrageService.
func NewBasicBarrageService() BarrageService {
	return &basicBarrageService{}
}

// New returns a BarrageService with all of the expected middleware wired in.
func New(middleware []Middleware) BarrageService {
	var svc BarrageService = NewBasicBarrageService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
