package endpoint

import (
	"context"
	service "github.com/GrantZheng/monolith_demo/sensitive_words/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CheckWordsRequest collects the request parameters for the CheckWords method.
type CheckWordsRequest struct {
	Words string `json:"words"`
}

// CheckWordsResponse collects the response parameters for the CheckWords method.
type CheckWordsResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCheckWordsEndpoint returns an endpoint that invokes CheckWords on the service.
func MakeCheckWordsEndpoint(s service.SensitiveWordsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckWordsRequest)
		rs, err := s.CheckWords(ctx, req.Words)
		return CheckWordsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CheckWordsResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CheckWords implements Service. Primarily useful in a client.
func (e Endpoints) CheckWords(ctx context.Context, words string) (rs string, err error) {
	request := CheckWordsRequest{Words: words}
	response, err := e.CheckWordsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CheckWordsResponse).Rs, response.(CheckWordsResponse).Err
}
