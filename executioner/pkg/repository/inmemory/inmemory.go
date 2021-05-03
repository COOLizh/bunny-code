// Package inmemory implements in-memory realization of SolutionsRepository interface using map and sync.RWMutex
package inmemory

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/greenteam1/executioner/pkg/pb"
)

// SolutionsRepository ...
type SolutionsRepository struct {
	data map[string]*pb.StatusHandleResponse
	m    sync.RWMutex
}

// NewSolutionsRepository ...
func NewSolutionsRepository() *SolutionsRepository {
	return &SolutionsRepository{
		data: make(map[string]*pb.StatusHandleResponse),
	}
}

// Add adds new element with id as key
func (r *SolutionsRepository) Add(_ context.Context, id string, sol *pb.StatusHandleResponse) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[id] = sol

	return nil
}

// Pop returns element by by id as key if it exists, and deletes it from storage
func (r *SolutionsRepository) Pop(_ context.Context, id string) (*pb.StatusHandleResponse, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	sol, ok := r.data[id]
	if !ok {
		return &pb.StatusHandleResponse{}, status.Error(
			codes.NotFound,
			fmt.Sprintf("no records with id %s", id),
		)
	}
	if ok {
		if !sol.Ready {
			return &pb.StatusHandleResponse{}, nil
		}
		delete(r.data, id)
	}

	return sol, nil
}
