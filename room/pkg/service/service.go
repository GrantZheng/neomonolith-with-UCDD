package service

import (
	"context"
)

// RoomService describes the service.
type RoomService interface {
	// Add your methods here
	Get(ctx context.Context, roomid string) (rs string, err error)
}

type basicRoomService struct{}

func (b *basicRoomService) Get(ctx context.Context, roomid string) (rs string, err error) {
	// TODO implement the business logic of Get
	return rs, err
}

// NewBasicRoomService returns a naive, stateless implementation of RoomService.
func NewBasicRoomService() RoomService {
	return &basicRoomService{}
}

// New returns a RoomService with all of the expected middleware wired in.
func New(middleware []Middleware) RoomService {
	var svc RoomService = NewBasicRoomService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
