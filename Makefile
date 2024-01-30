.PHONY: build test server client

build:
	@docker build -f Dockerfile.server -t tcp_pow_server --target runtime . \
	&& docker build -f Dockerfile.client -t tcp_pow_client --target runtime .

test:
	@go test -v ./...

server:
	@docker run -d --rm -p 9999:9999 -e SERVER_PORT=9999 -e POW_COMPLEXITY=3 --name local_tcp_pow_server tcp_pow_server


stop:
	@docker stop local_tcp_pow_server

client:
	@SERVER_ADDR=localhost SERVER_PORT=9999 ROW_COMPLEXITY=3 go run cmd/client/main.go
