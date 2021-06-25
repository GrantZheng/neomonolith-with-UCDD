package service

import (
	"context"
)

// SensitiveWordsService describes the service.
type SensitiveWordsService interface {
	// Add your methods here
	CheckWords(ctx context.Context, words string) (rs string, err error)
}

type basicSensitiveWordsService struct{}

func (b *basicSensitiveWordsService) CheckWords(ctx context.Context, words string) (rs string, err error) {
	// TODO implement the business logic of CheckWords
	return rs, err
}

// NewBasicSensitiveWordsService returns a naive, stateless implementation of SensitiveWordsService.
func NewBasicSensitiveWordsService() SensitiveWordsService {
	return &basicSensitiveWordsService{}
}

// New returns a SensitiveWordsService with all of the expected middleware wired in.
func New(middleware []Middleware) SensitiveWordsService {
	var svc SensitiveWordsService = NewBasicSensitiveWordsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
