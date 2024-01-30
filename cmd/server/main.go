package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tcp_pow/internal/pow"
	"tcp_pow/pkg/tcp"
	"time"
)

func main() {
	// Read environment variables
	port := getEnv("SERVER_PORT", "9999")                                       // Default port is 9999
	timeoutMillis, err := strconv.Atoi(getEnv("SERVER_TIMEOUT_MILLIS", "5000")) // Default timeout is 5000ms
	if err != nil {
		log.Fatalf("Invalid timeout value: %v", err)
	}

	powComplexity, err := strconv.Atoi(getEnv("POW_COMPLEXITY", "3")) // Default PoW complexity is 3
	if err != nil {
		log.Fatalf("Invalid PoW complexity value: %v", err)
	}

	// Initialize the PoW validator
	powValidator := pow.NewSha256PoW(powComplexity)

	// Initialize and start the TCP server
	serverAddress := fmt.Sprintf("0.0.0.0:%s", port)
	server := tcp.NewTCPServer(serverAddress, powValidator, time.Duration(timeoutMillis)*time.Millisecond)
	fmt.Printf("Starting server on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
