package services

import (
	"context"

	"github.com/Evertras/sandbox-grpc-js/lib/messages"
)

// EchoService implements the Echo service
type EchoService struct{}

// SayHello just spits the message back to the client
func (e *EchoService) SayHello(ctx context.Context, msg *messages.EchoMessage) (*messages.EchoMessage, error) {
	return msg, nil
}
