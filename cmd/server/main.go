package main

import (
	"flag"
	"log"
	"net"

	"github.com/Evertras/sandbox-grpc-js/lib"
	"github.com/Evertras/sandbox-grpc-js/lib/data"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	data.RegisterEchoServer(grpcServer, &lib.EchoService{})
	grpcServer.Serve(lis)
}
