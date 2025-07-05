package server

import (
	"context"
	"fmt"

	"kvstore-grpc/gen/kvstorepb"
)

type InMemoryKVStore struct {
	kvstorepb.UnimplementedKVStoreServer // Required for forward compatibility
	store                                map[string]string
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return &InMemoryKVStore{
		store: make(map[string]string),
	}
}

func (s *InMemoryKVStore) Set(ctx context.Context, req *kvstorepb.SetRequest) (*kvstorepb.Empty, error) {
	s.store[req.Key] = req.Value
	return &kvstorepb.Empty{}, nil
}

func (s *InMemoryKVStore) Get(ctx context.Context, req *kvstorepb.GetRequest) (*kvstorepb.GetResponse, error) {
	// Get key from store
	value, found := s.store[req.Key]
	// Check if key exists in store
	if !found {
		return nil, fmt.Errorf("key %v not found", req.Key)
	}
	return &kvstorepb.GetResponse{
		Value: value,
	}, nil
}

func (s *InMemoryKVStore) Delete(ctx context.Context, req *kvstorepb.DeleteRequest) (*kvstorepb.Empty, error) {
	// Check if key exists in store
	_, found := s.store[req.Key]
	if !found {
		return nil, fmt.Errorf("key %v not found", req.Key)
	}
	delete(s.store, req.Key)
	return &kvstorepb.Empty{}, nil
}
