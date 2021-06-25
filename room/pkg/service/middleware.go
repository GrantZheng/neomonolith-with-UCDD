package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(RoomService) RoomService

type loggingMiddleware struct {
	logger log.Logger
	next   RoomService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a RoomService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next RoomService) RoomService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context, roomid string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Get", "roomid", roomid, "rs", rs, "err", err)
	}()
	return l.next.Get(ctx, roomid)
}
