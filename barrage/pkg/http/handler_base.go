package http

import (
	http1 "net/http"

	endpoint "github.com/GrantZheng/monolith_demo/barrage/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
)

// InitHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func InitHTTPHandler(m *http1.ServeMux, endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	makeSendHandler(m, endpoints, options["Send"])
	return m
}
