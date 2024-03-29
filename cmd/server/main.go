package main

import (
	"fmt"
	"log"
	"tcp_pow/internal/pow"
	"tcp_pow/pkg/config"
	"tcp_pow/pkg/tcp"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	powValidator := pow.NewSha256PoW(cfg.PowComplexity)

	serverAddress := fmt.Sprintf("%s:%s", cfg.ServerAddr, cfg.ServerPort)
	server := tcp.NewTCPServer(serverAddress, powValidator, time.Duration(cfg.TimeoutMillis)*time.Millisecond)
	fmt.Printf("Starting server %s:%s", cfg.ServerAddr, cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
