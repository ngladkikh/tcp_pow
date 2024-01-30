package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"tcp_pow/internal/pow"
	"tcp_pow/internal/quotes"
	"time"

	"github.com/google/uuid"
)

type TCPServer struct {
	addr         string
	powValidator pow.PoW
	timeout      time.Duration
}

func NewTCPServer(addr string, validator pow.PoW, timeout time.Duration) *TCPServer {
	return &TCPServer{
		addr:         addr,
		powValidator: validator,
		timeout:      timeout,
	}
}

func (s *TCPServer) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// Handle error, possibly continue to the next loop iteration
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	challenge := uuid.New().String()

	fmt.Fprintln(conn, challenge)

	conn.SetReadDeadline(time.Now().Add(s.timeout))
	nonce, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	nonce = strings.TrimSpace(nonce)

	if s.powValidator.Validate(challenge, nonce) {
		fmt.Fprintln(conn, quotes.Quote())
	}
}
