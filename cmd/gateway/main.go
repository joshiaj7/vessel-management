package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joshiaj7/vessel-management/internal/config"
)

func main() {
	cfg, err := config.NewGatewayServer()
	if err != nil {
		log.Fatalf("Load Gateway Server Failed: %v", err)
		return
	}

	fmt.Println("Listening to port 8080")
	http.ListenAndServe(":8080", cfg.Router)
}
