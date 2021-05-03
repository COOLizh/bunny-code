package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"

	"gitlab.com/greenteam1/task_repo/pkg/client"
	"gitlab.com/greenteam1/task_repo/pkg/models"

	"gitlab.com/greenteam1/task_repo/pkg/repository"
)

// SolutionService ...
type SolutionService struct {
	SolutionRepository repository.SolutionRepositoryInterface
}

// NewSolutionService returns created SolutionService struct
func NewSolutionService(sr repository.SolutionRepositoryInterface) *SolutionService {
	return &SolutionService{
		SolutionRepository: sr,
	}
}

// StoreSolution ...
func (ss *SolutionService) StoreSolution(ctx context.Context, uuid string, userID int, taskID int, lang string, sol []byte) (time.Time, error) {
	return ss.SolutionRepository.StoreSolution(ctx, uuid, lang, userID, taskID, sol)
}

// GetSolutionResult ...
func (ss *SolutionService) GetSolutionResult(ctx context.Context, exAddr, uuid string) (models.SolutionResult, bool, error) {
	resDB, tcCount, err := ss.SolutionRepository.GetResultAndTestCaseCountByID(ctx, uuid)
	if err == pgx.ErrNoRows {
		resCli, ready, errCli := client.GetSolutionResult(ctx, exAddr, uuid)
		if !ready || errCli != nil {
			return models.SolutionResult{}, false, errCli
		}

		resCli.TestsCount = tcCount
		errCli = ss.SolutionRepository.StoreResult(ctx, uuid, resCli)
		if errCli != nil {
			return models.SolutionResult{}, false, errCli
		}

		return resCli, ready, nil
	}
	if err != nil {
		return models.SolutionResult{}, false, err
	}

	resDB.TestsCount = tcCount
	return resDB, true, nil
}

// GetSolutionsHistory ...
func (ss *SolutionService) GetSolutionsHistory(ctx context.Context, userID, taskID int) ([]models.SolutionHistoryItem, error) {
	return ss.SolutionRepository.GetSolutionsHistory(ctx, userID, taskID)
}
