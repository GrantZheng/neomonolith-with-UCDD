package service

import "context"

// LiveRoomService describes the service.
type LiveRoomService interface {
	// Add your methods here
	Get(ctx context.Context, roomid string) (rs string, err error)
	Open(ctx context.Context, roomid string) (err error)
	Close(ctx context.Context, roomid string) (err error)
}

type basicLiveRoomService struct{}

func (b *basicLiveRoomService) Get(ctx context.Context, roomid string) (rs string, err error) {
	// TODO implement the business logic of Get
	return rs, err
}
func (b *basicLiveRoomService) Open(ctx context.Context, roomid string) (err error) {
	// TODO implement the business logic of Open
	return err
}
func (b *basicLiveRoomService) Close(ctx context.Context, roomid string) (err error) {
	// TODO implement the business logic of Close
	return err
}

// NewBasicLiveRoomService returns a naive, stateless implementation of LiveRoomService.
func NewBasicLiveRoomService() LiveRoomService {
	return &basicLiveRoomService{}
}

// New returns a LiveRoomService with all of the expected middleware wired in.
func New(middleware []Middleware) LiveRoomService {
	var svc LiveRoomService = NewBasicLiveRoomService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
