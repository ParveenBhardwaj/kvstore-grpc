# Variables
BINARY_NAME=kvstore-grpc
SERVER_MAIN=./cmd/server/main.go
CLIENT_MAIN=./cmd/client/client.go

# build server binary
build:
	go build -o $(BINARY_NAME) $(SERVER_MAIN)

# run server
run-server:
	go run $(SERVER_MAIN)

# run client
run-client:
	go run $(CLIENT_MAIN)

# run unit tests
test:
	go test ./internal/... -v

# regenerate gRPC code from proto
proto-gen:
	protoc \
  	--proto_path=proto \
  	--go_out=gen \
  	--go-grpc_out=gen \
  	proto/kvstore.proto

# build docker image
docker-build:
	docker build -t $(BINARY_NAME) .

# run docker container on port 50051
docker-run
	docker run -p 50051:50051 $(BINARY_NAME)

lint:
	golangci-lint run

clean:
	go clean
	rm -f $(BINARY_NAME)