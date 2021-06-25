package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(WalletService) WalletService

type loggingMiddleware struct {
	logger log.Logger
	next   WalletService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a WalletService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next WalletService) WalletService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Deduct(ctx context.Context, uid string, money int) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Deduct", "uid", uid, "money", money, "rs", rs, "err", err)
	}()
	return l.next.Deduct(ctx, uid, money)
}
