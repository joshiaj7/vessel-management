package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joshiaj7/vessel-management/internal/config"
)

func main() {
	gatewayServer, err := config.NewGatewayServer()
	if err != nil {
		log.Fatalf("failed to create new gateway: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(server *config.GatewayServer) {
		log.Printf("Starting Gateway on: %v\n", gatewayServer.Host)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error listening gateway: %v", err)
		}
	}(gatewayServer)

	<-sigChan

	log.Println("Shutting down the gateway Server...")
	err = gatewayServer.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}
	log.Println("gateway Server gracefully stopped")
}
