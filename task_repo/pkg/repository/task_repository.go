package repository

import (
	"context"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// TaskRepository ...
type TaskRepository struct {
	Task TaskRepositoryInterface
}

// NewTaskRepository returns new TaskRepository
func NewTaskRepository(db TaskRepositoryInterface) *TaskRepository {
	return &TaskRepository{
		Task: db,
	}
}

// TaskRepositoryInterface represent the task's repository contract
type TaskRepositoryInterface interface {
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id int) (models.Task, error)
	GetTestCasesByTaskID(ctx context.Context, taskID int) ([]models.TestCase, error)
}
