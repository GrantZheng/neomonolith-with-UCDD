package service

import (
	"context"
	"errors"

	roomsvc "github.com/GrantZheng/monolith_demo/room/pkg/service"
)

type roomSvc struct{}

func (r *roomSvc) InitRoomSvcHandler(_ context.Context) roomsvc.RoomService {
	// When room service move to a microservice that deploys independently.
	// You should replace the roomsvc.NewBasicRoomService() with New(instance string, options map[string][]http.ClientOption)
	// in github.com/GrantZheng/monolith_demo/room/client/http/http.go
	roomHandler := roomsvc.NewBasicRoomService()
	return roomHandler
}

func (r *roomSvc) GetRoom(ctx context.Context, roomid string) (string, error) {
	roomHandler := r.InitRoomSvcHandler(ctx)
	if roomHandler == nil {
		return "", errors.New("room svc error")
	}

	room, err := roomHandler.Get(ctx, roomid)
	if err != nil {
		return "", err
	}

	return room, err
}
