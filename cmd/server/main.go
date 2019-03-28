package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/Evertras/sandbox-grpc-js/lib/messages"
	"github.com/Evertras/sandbox-grpc-js/lib/services"
)

func runHTTP(ctx context.Context) error {
	echoService := &services.EchoService{}

	echoHandler := messages.NewEchoServer(echoService, nil)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./front")))
	mux.Handle(messages.EchoPathPrefix, echoHandler)

	server := http.Server{
		Handler: mux,
		Addr:    "localhost:8000",
	}

	go func() {
		<-ctx.Done()

		server.Close()
	}()

	log.Println("Running HTTP server")

	return server.ListenAndServe()
}

func main() {
	c := make(chan os.Signal, 1)

	// Quit on any signal, for simplicity/safety
	signal.Notify(c)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

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
