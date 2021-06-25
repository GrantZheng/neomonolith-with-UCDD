package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "user/pkg/service"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	S string `json:"s"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		rs, err := s.Add(ctx, req.S)
		return AddResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, s string) (rs string, err error) {
	request := AddRequest{S: s}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).Rs, response.(AddResponse).Err
}
