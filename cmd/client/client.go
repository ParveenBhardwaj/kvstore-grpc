package main

import (
	"context"
	"fmt"
	"log"

	"kvstore-grpc/gen/kvstorepb"

	"google.golang.org/grpc"
)

func main() {
	// Set up connection to the server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := kvstorepb.NewKVStoreClient(conn)

	// Perform set operation
	_, err = client.Set(context.Background(), &kvstorepb.SetRequest{Key: "hello", Value: "world"})
	if err != nil {
		log.Fatalf("Set failed: %v", err)
	}

	// Perform Get operation
	resp, err := client.Get(context.Background(), &kvstorepb.GetRequest{Key: "hello"})
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	fmt.Printf("Get response %s\n", resp.Value)

}
