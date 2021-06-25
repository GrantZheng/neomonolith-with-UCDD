package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/wallet/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// DeductRequest collects the request parameters for the Deduct method.
type DeductRequest struct {
	Uid   string `json:"uid"`
	Money int    `json:"money"`
}

// DeductResponse collects the response parameters for the Deduct method.
type DeductResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeductEndpoint returns an endpoint that invokes Deduct on the service.
func MakeDeductEndpoint(s service.WalletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeductRequest)
		rs, err := s.Deduct(ctx, req.Uid, req.Money)
		return DeductResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeductResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Deduct implements Service. Primarily useful in a client.
func (e Endpoints) Deduct(ctx context.Context, uid string, money int) (rs string, err error) {
	request := DeductRequest{
		Money: money,
		Uid:   uid,
	}
	response, err := e.DeductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeductResponse).Rs, response.(DeductResponse).Err
}
