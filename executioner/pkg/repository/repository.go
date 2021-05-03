// Package repository ...
package repository

import (
	"context"

	"gitlab.com/greenteam1/executioner/pkg/pb"
)

// SolutionsRepository describes storage for completed tasks
type SolutionsRepository interface {
	Add(context.Context, string, *pb.StatusHandleResponse) error
	Pop(context.Context, string) (*pb.StatusHandleResponse, error)
}
