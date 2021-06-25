package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(GiftService) GiftService

type loggingMiddleware struct {
	logger log.Logger
	next   GiftService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GiftService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GiftService) GiftService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Give(ctx context.Context, from string, tu string, gift_id string, gift_num string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Give", "from", from, "tu", tu, "gift_id", gift_id, "gift_num", gift_num, "rs", rs, "err", err)
	}()
	return l.next.Give(ctx, from, tu, gift_id, gift_num)
}
