package service

import (
	"context"
)

// VideoContentService describes the service.
type VideoContentService interface {
	// Add your methods here
	GetVideoAddr(ctx context.Context, roomid string) (rs string, err error)
}

type basicVideoContentService struct{}

func (b *basicVideoContentService) GetVideoAddr(ctx context.Context, roomid string) (rs string, err error) {
	// TODO implement the business logic of GetVideoAddr
	return rs, err
}

// NewBasicVideoContentService returns a naive, stateless implementation of VideoContentService.
func NewBasicVideoContentService() VideoContentService {
	return &basicVideoContentService{}
}

// New returns a VideoContentService with all of the expected middleware wired in.
func New(middleware []Middleware) VideoContentService {
	var svc VideoContentService = NewBasicVideoContentService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
