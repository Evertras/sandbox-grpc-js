package lib

import (
	"context"

	"github.com/Evertras/sandbox-grpc-js/lib/data"
)

// EchoService implements the Echo service
type EchoService struct{}

// SayHello just spits the message back to the client
func (e *EchoService) SayHello(ctx context.Context, msg *data.EchoMessage) (*data.EchoMessage, error) {
	return msg, nil
}
