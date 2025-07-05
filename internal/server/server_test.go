package server

import (
	"context"
	"testing"

	"kvstore-grpc/gen/kvstorepb"
)

func TestSetStoreValue(t *testing.T) {
	// Create a new store
	s := NewInMemoryKVStore()

	req := &kvstorepb.SetRequest{
		Key:   "hello",
		Value: "world",
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

func TestGetStoreValue(t *testing.T) {
	//  Create a new store
	s := NewInMemoryKVStore()
	req := &kvstorepb.SetRequest{
		Key:   "hello",
		Value: "world",
	}

	_, err := s.Set(context.Background(), req)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	getReq := &kvstorepb.GetRequest{
		Key: "hello",
	}

	got, err := s.Get(context.Background(), getReq)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	want := "world"

	if got.Value != want {
		t.Errorf("expected value %q, got %q", want, got)
	}

}

func TestDeleteStoreValue(t *testing.T) {
	// Create a new store
	s := NewInMemoryKVStore()
	setReq := &kvstorepb.SetRequest{
		Key:   "hello",
		Value: "world",
	}

	_, err := s.Set(context.Background(), setReq)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	deleteReq := &kvstorepb.DeleteRequest{
		Key: "hello",
	}
	_, err = s.Delete(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	if len(s.store) != 0 {
		t.Fatalf("Expected store to be empty, but store lenght is %v", len(s.store))
	}

}
