package service

import (
	"context"
)

// MessageChannelService describes the service.
type MessageChannelService interface {
	// Add your methods here
	Send(ctx context.Context,msg string, channel_id string)(err error)
}
