package service

import (
	"context"
)

// MessageChannelService describes the service.
type MessageChannelService interface {
	// Add your methods here
	Send(ctx context.Context, msg string, channel_id string) (err error)
}

type basicMessageChannelService struct{}

func (b *basicMessageChannelService) Send(ctx context.Context, msg string, channel_id string) (err error) {
	// TODO implement the business logic of Send
	return err
}

// NewBasicMessageChannelService returns a naive, stateless implementation of MessageChannelService.
func NewBasicMessageChannelService() MessageChannelService {
	return &basicMessageChannelService{}
}

// New returns a MessageChannelService with all of the expected middleware wired in.
func New(middleware []Middleware) MessageChannelService {
	var svc MessageChannelService = NewBasicMessageChannelService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
