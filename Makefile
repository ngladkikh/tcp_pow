.PHONY: test server client build integration

test:
	@go test -v ./...

server:
	@SERVER_ADDR=localhost SERVER_PORT=9999 ROW_COMPLEXITY=3 go run cmd/server/main.go

client:
	@SERVER_ADDR=localhost SERVER_PORT=9999 ROW_COMPLEXITY=3 go run cmd/client/main.go

build:
	@docker build -f Dockerfile.server -t tcp_pow_server --target runtime . \
	&& docker build -f Dockerfile.client -t tcp_pow_client --target runtime .

integration:
	docker-compose up --abort-on-container-exit --exit-code-from client
	docker-compose down