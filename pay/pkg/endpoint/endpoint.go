package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/pay/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// PayRequest collects the request parameters for the Pay method.
type PayRequest struct {
	Account float32 `json:"account"`
	UserId  string  `json:"user_id"`
}

// PayResponse collects the response parameters for the Pay method.
type PayResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakePayEndpoint returns an endpoint that invokes Pay on the service.
func MakePayEndpoint(s service.PayService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PayRequest)
		rs, err := s.Pay(ctx, req.Account, req.UserId)
		return PayResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r PayResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Pay implements Service. Primarily useful in a client.
func (e Endpoints) Pay(ctx context.Context, account float32, user_id string) (rs string, err error) {
	request := PayRequest{
		Account: account,
		UserId:  user_id,
	}
	response, err := e.PayEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PayResponse).Rs, response.(PayResponse).Err
}
