package main

import (
	"fmt"
	"log"
	"net"

	"kvstore-grpc/gen/kvstorepb"
	"kvstore-grpc/internal/server"

	"google.golang.org/grpc"
)

func main() {
	// Setup the server
	grpcServer := grpc.NewServer()
	// Initialise memory store
	kvStoreServer := server.NewInMemoryKVStore()

	// Register the server with gRPC
	kvstorepb.RegisterKVStoreServer(grpcServer, kvStoreServer)

	// Set up the listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Start the server
	fmt.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
