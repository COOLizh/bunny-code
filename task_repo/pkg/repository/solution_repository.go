// Package repository ...
package repository

import (
	"context"
	"time"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// SolutionRepository ...
type SolutionRepository struct {
	Solution SolutionRepositoryInterface
}

// NewSolutionRepository returns new UserRepository
func NewSolutionRepository(db SolutionRepositoryInterface) *SolutionRepository {
	return &SolutionRepository{
		Solution: db,
	}
}

// SolutionRepositoryInterface represent the user's repository contract
type SolutionRepositoryInterface interface {
	StoreSolution(context.Context, string, string, int, int, []byte) (time.Time, error)
	StoreResult(context.Context, string, models.SolutionResult) error
	GetResultByID(context.Context, string) (models.SolutionResult, error)
	GetResultAndTestCaseCountByID(context.Context, string) (models.SolutionResult, int64, error)
	GetSolutionsHistory(ctx context.Context, userID, taskID int) ([]models.SolutionHistoryItem, error)
}
