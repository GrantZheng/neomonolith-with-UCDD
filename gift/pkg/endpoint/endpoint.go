package endpoint

import (
	"context"
	service "gift/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// ListRequest collects the request parameters for the List method.
type ListRequest struct {
	S string `json:"s"`
}

// ListResponse collects the response parameters for the List method.
type ListResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeListEndpoint returns an endpoint that invokes List on the service.
func MakeListEndpoint(s service.GiftService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		rs, err := s.List(ctx, req.S)
		return ListResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ListResponse) Failed() error {
	return r.Err
}

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	GiftId string `json:"gift_id"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.GiftService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		rs, err := s.Send(ctx, req.From, req.To, req.GiftId)
		return SendResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SendResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// List implements Service. Primarily useful in a client.
func (e Endpoints) List(ctx context.Context, s string) (rs string, err error) {
	request := ListRequest{S: s}
	response, err := e.ListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListResponse).Rs, response.(ListResponse).Err
}

// Send implements Service. Primarily useful in a client.
func (e Endpoints) Send(ctx context.Context, from string, to string, gift_id string) (rs string, err error) {
	request := SendRequest{
		From:   from,
		GiftId: gift_id,
		To:     to,
	}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).Rs, response.(SendResponse).Err
}
