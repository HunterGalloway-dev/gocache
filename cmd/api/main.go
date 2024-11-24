package main

import (
	"context"
	"gocache/internal/logger"
	"gocache/internal/server"

	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	logger.Logger.Warn("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		logger.Logger.Fatalf("could not gracefully shutdown the server: %v\n", err)
	}

	//log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {

	server, err := server.NewServer()
	if err != nil {
		logger.Logger.Fatalf("could not create server: %v\n", err)
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Logger.Fatalf("could not listen on %s: %v\n", server.Addr, err)
	}

	// Wait for the graceful shutdown to complete
	<-done
	logger.Logger.Info("server shutdown complete")
}
