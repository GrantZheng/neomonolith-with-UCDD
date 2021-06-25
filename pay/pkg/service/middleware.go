package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(PayService) PayService

type loggingMiddleware struct {
	logger log.Logger
	next   PayService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PayService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PayService) PayService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Pay(ctx context.Context, account float32, user_id string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Pay", "account", account, "user_id", user_id, "rs", rs, "err", err)
	}()
	return l.next.Pay(ctx, account, user_id)
}
