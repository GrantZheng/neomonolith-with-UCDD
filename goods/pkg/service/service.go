package service

import "context"

// GoodsService describes the service.
type GoodsService interface {
	// Add your methods here
	Get(ctx context.Context, id string) (rs string, err error)
}

type basicGoodsService struct{}

func (b *basicGoodsService) Get(ctx context.Context, id string) (rs string, err error) {
	// TODO implement the business logic of Get
	return rs, err
}

// NewBasicGoodsService returns a naive, stateless implementation of GoodsService.
func NewBasicGoodsService() GoodsService {
	return &basicGoodsService{}
}

// New returns a GoodsService with all of the expected middleware wired in.
func New(middleware []Middleware) GoodsService {
	var svc GoodsService = NewBasicGoodsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
