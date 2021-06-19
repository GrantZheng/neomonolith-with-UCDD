package endpoint

import (
	service "barrage/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.BarrageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		rs, err := s.Send(ctx, req.From, req.To, req.Msg)
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

// Send implements Service. Primarily useful in a client.
func (e Endpoints) Send(ctx context.Context, from string, to string, msg string) (rs string, err error) {
	request := SendRequest{
		From: from,
		Msg:  msg,
		To:   to,
	}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).Rs, response.(SendResponse).Err
}
