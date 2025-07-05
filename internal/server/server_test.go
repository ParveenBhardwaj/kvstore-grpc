package server

import (
	"context"
	"testing"

	"kvstore-grpc/gen/kvstorepb"
)

func TestSetStoreValue(t *testing.T) {
	s := NewInMemoryKVStore()

	req := &kvstorepb.SetRequest{
		Key:   "hello",
		value: "world",
	}

	_, err := s.Set(context.Background(), req)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// We check internal state to confirm it's stored
	got := s.store["hello"]
	want := "world"

	if got != want {
		t.Errorf("expected value %q, got %q", want, got)
	}
}
