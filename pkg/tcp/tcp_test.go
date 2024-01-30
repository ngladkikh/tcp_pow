package tcp

import (
	"bufio"
	"fmt"
	"net"
	"tcp_pow/internal/pow"
	"testing"
	"time"
)

const ServerAddr = "localhost:54321"

func TestTCPServer(t *testing.T) {
	stubPow := pow.NewStubPow("foo")
	server := NewTCPServer(ServerAddr, stubPow, 100*time.Millisecond)
	go server.ListenAndServe()

	t.Run("Valid Nonce and Quote Returned", func(t *testing.T) {
		conn, err := net.Dial("tcp", ServerAddr)
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)
		challenge, _ := reader.ReadString('\n')
		if challenge == "" {
			t.Error("Expected a challenge, got empty response")
		}

		fmt.Fprintln(conn, "foo")

		quote, _ := reader.ReadString('\n')
		if quote == "" {
			t.Error("Expected a quote, got empty response")
		}
	})

	t.Run("Invalid Nonce and Connection is Closed", func(t *testing.T) {
		conn, err := net.Dial("tcp", ServerAddr)
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)
		_, _ = reader.ReadString('\n')

		// Send an invalid nonce
		fmt.Fprintln(conn, "invalid_nonce")

		// Try to read from the connection
		_, err = reader.ReadString('\n')
		if err == nil {
			t.Error("Expected connection to be closed, but it was not")
		}
	})

}
