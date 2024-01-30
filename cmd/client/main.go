package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"tcp_pow/internal/pow"
	"tcp_pow/pkg/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	fullServerAddr := fmt.Sprintf("%s:%s", cfg.ServerAddr, cfg.ServerPort)
	conn, err := net.Dial("tcp", fullServerAddr)
	if err != nil {
		fmt.Printf("Failed to connect to server: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Receive PoW challenge
	challenge, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read challenge: %s\n", err)
		os.Exit(1)
	}
	challenge = strings.TrimSpace(challenge)

	fmt.Printf("Received challenge: %s\n", challenge)

	// Brute force to find valid nonce
	nonce := findValidNonce(challenge, cfg.PowComplexity)
	fmt.Printf("Found valid nonce: %s\n", nonce)

	// Send nonce back to server
	fmt.Fprintln(conn, nonce)

	// Receive quote
	quote, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read quote: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Received quote: %s\n", quote)
}

func findValidNonce(challenge string, complexity int) string {
	powValidator := pow.NewSha256PoW(complexity)
	nonce := 0
	for {
		nonceStr := fmt.Sprint(nonce)
		if powValidator.Validate(challenge, nonceStr) {
			return nonceStr
		}
		nonce++
	}
}
