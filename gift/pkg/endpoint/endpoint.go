package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/gift/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GiveRequest collects the request parameters for the Give method.
type GiveRequest struct {
	From    string `json:"from"`
	Tu      string `json:"tu"`
	GiftId  string `json:"gift_id"`
	GiftNum string `json:"gift_num"`
}

// GiveResponse collects the response parameters for the Give method.
type GiveResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGiveEndpoint returns an endpoint that invokes Give on the service.
func MakeGiveEndpoint(s service.GiftService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GiveRequest)
		rs, err := s.Give(ctx, req.From, req.Tu, req.GiftId, req.GiftNum)
		return GiveResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GiveResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Give implements Service. Primarily useful in a client.
func (e Endpoints) Give(ctx context.Context, from string, tu string, gift_id string, gift_num string) (rs string, err error) {
	request := GiveRequest{
		From:    from,
		GiftId:  gift_id,
		GiftNum: gift_num,
		Tu:      tu,
	}
	response, err := e.GiveEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GiveResponse).Rs, response.(GiveResponse).Err
}
