package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/GrantZheng/monolith_demo/live_room/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeGetHandler creates the handler logic
func makeGetHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get", http1.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...))
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeOpenHandler creates the handler logic
func makeOpenHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/open", http1.NewServer(endpoints.OpenEndpoint, decodeOpenRequest, encodeOpenResponse, options...))
}

// decodeOpenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeOpenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.OpenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeOpenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeOpenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCloseHandler creates the handler logic
func makeCloseHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/close", http1.NewServer(endpoints.CloseEndpoint, decodeCloseRequest, encodeCloseResponse, options...))
}

// decodeCloseRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCloseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CloseRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCloseResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCloseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
