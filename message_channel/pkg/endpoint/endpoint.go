package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/message_channel/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	Msg       string `json:"msg"`
	ChannelId string `json:"channel_id"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	Err error `json:"err"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.MessageChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		err := s.Send(ctx, req.Msg, req.ChannelId)
		return SendResponse{Err: err}, nil
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
func (e Endpoints) Send(ctx context.Context, msg string, channel_id string) (err error) {
	request := SendRequest{
		ChannelId: channel_id,
		Msg:       msg,
	}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).Err
}
