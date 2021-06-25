package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(GoodsService) GoodsService

type loggingMiddleware struct {
	logger log.Logger
	next   GoodsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GoodsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GoodsService) GoodsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Get", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Get(ctx, s)
}
