package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(VideoContentService) VideoContentService

type loggingMiddleware struct {
	logger log.Logger
	next   VideoContentService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a VideoContentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next VideoContentService) VideoContentService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetVideoAddr(ctx context.Context, roomid string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "GetVideoAddr", "roomid", roomid, "rs", rs, "err", err)
	}()
	return l.next.GetVideoAddr(ctx, roomid)
}
