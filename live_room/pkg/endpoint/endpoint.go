package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/live_room/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Roomid string `json:"roomid"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.LiveRoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		rs, err := s.Get(ctx, req.Roomid)
		return GetResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// OpenRequest collects the request parameters for the Open method.
type OpenRequest struct {
	Roomid string `json:"roomid"`
}

// OpenResponse collects the response parameters for the Open method.
type OpenResponse struct {
	Err error `json:"err"`
}

// MakeOpenEndpoint returns an endpoint that invokes Open on the service.
func MakeOpenEndpoint(s service.LiveRoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenRequest)
		err := s.Open(ctx, req.Roomid)
		return OpenResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r OpenResponse) Failed() error {
	return r.Err
}

// CloseRequest collects the request parameters for the Close method.
type CloseRequest struct {
	Roomid string `json:"roomid"`
}

// CloseResponse collects the response parameters for the Close method.
type CloseResponse struct {
	Err error `json:"err"`
}

// MakeCloseEndpoint returns an endpoint that invokes Close on the service.
func MakeCloseEndpoint(s service.LiveRoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CloseRequest)
		err := s.Close(ctx, req.Roomid)
		return CloseResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r CloseResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, roomid string) (rs string, err error) {
	request := GetRequest{Roomid: roomid}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Rs, response.(GetResponse).Err
}

// Open implements Service. Primarily useful in a client.
func (e Endpoints) Open(ctx context.Context, roomid string) (err error) {
	request := OpenRequest{Roomid: roomid}
	response, err := e.OpenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(OpenResponse).Err
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close(ctx context.Context, roomid string) (err error) {
	request := CloseRequest{Roomid: roomid}
	response, err := e.CloseEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CloseResponse).Err
}
