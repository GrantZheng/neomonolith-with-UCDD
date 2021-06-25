package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(MessageChannelService) MessageChannelService

type loggingMiddleware struct {
	logger log.Logger
	next   MessageChannelService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a MessageChannelService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MessageChannelService) MessageChannelService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Send(ctx context.Context, msg string, channel_id string) (err error) {
	defer func() {
		l.logger.Log("method", "Send", "msg", msg, "channel_id", channel_id, "err", err)
	}()
	return l.next.Send(ctx, msg, channel_id)
}
