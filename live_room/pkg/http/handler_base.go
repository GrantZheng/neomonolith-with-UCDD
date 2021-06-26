package http

import (
	http1 "net/http"

	endpoint "github.com/GrantZheng/monolith_demo/live_room/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
)

// InitHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func InitHTTPHandler(m *http1.ServeMux, endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	makeGetHandler(m, endpoints, options["Get"])
	makeOpenHandler(m, endpoints, options["Open"])
	makeCloseHandler(m, endpoints, options["Close"])
	return m
}
