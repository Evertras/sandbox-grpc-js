package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/Evertras/sandbox-grpc-js/lib/services"
)

func runRPC(ctx context.Context) error {
	lis, err := net.Listen("tcp", "localhost:8001")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := services.NewServer()

	log.Println("Running")

	go func() {
		<-ctx.Done()

		server.GracefulStop()
	}()

	return server.Serve(lis)
}

func runHTTP(ctx context.Context) error {
	server := http.Server{
		// Quick and dirty, exposes source... clean up for reals if it's for reals
		Handler: http.FileServer(http.Dir("./front")),
		Addr:    "localhost:8000",
	}

	go func() {
		<-ctx.Done()

		server.Close()
	}()

	return server.ListenAndServe()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		wg.Add(1)
		err := runRPC(ctx)

		if err != nil {
			log.Println("RPC server failed:", err)
		}

		wg.Done()
	}()

	go func() {
		wg.Add(1)
		err := runHTTP(ctx)

		if err != nil && err.Error() != "http: Server closed" {
			log.Println("HTTP server failed:", err)
		}

		wg.Done()
	}()

	<-c

	cancel()

	wg.Wait()

	log.Println("Graceful shutdown complete")
}
