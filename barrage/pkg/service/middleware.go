package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BarrageService) BarrageService

type loggingMiddleware struct {
	logger log.Logger
	next   BarrageService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BarrageService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BarrageService) BarrageService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Send(ctx context.Context, from string, to string, msg string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Send", "from", from, "to", to, "msg", msg, "rs", rs, "err", err)
	}()
	return l.next.Send(ctx, from, to, msg)
}
