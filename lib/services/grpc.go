package services

import (
	"github.com/Evertras/sandbox-grpc-js/lib/data"
	"google.golang.org/grpc"
)

// NewServer returns a new GRPC server with all our services registered
func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	// Add registrations here
	data.RegisterEchoServer(grpcServer, &EchoService{})

	return grpcServer
}
