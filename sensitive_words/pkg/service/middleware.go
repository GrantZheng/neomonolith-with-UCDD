package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SensitiveWordsService) SensitiveWordsService

type loggingMiddleware struct {
	logger log.Logger
	next   SensitiveWordsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SensitiveWordsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SensitiveWordsService) SensitiveWordsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CheckWords(ctx context.Context, words string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CheckWords", "words", words, "rs", rs, "err", err)
	}()
	return l.next.CheckWords(ctx, words)
}
