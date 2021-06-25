package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/video_content/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetVideoAddrRequest collects the request parameters for the GetVideoAddr method.
type GetVideoAddrRequest struct {
	Roomid string `json:"roomid"`
}

// GetVideoAddrResponse collects the response parameters for the GetVideoAddr method.
type GetVideoAddrResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetVideoAddrEndpoint returns an endpoint that invokes GetVideoAddr on the service.
func MakeGetVideoAddrEndpoint(s service.VideoContentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetVideoAddrRequest)
		rs, err := s.GetVideoAddr(ctx, req.Roomid)
		return GetVideoAddrResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetVideoAddrResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetVideoAddr implements Service. Primarily useful in a client.
func (e Endpoints) GetVideoAddr(ctx context.Context, roomid string) (rs string, err error) {
	request := GetVideoAddrRequest{Roomid: roomid}
	response, err := e.GetVideoAddrEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetVideoAddrResponse).Rs, response.(GetVideoAddrResponse).Err
}
