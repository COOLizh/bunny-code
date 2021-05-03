package service

import (
	"context"

	"gitlab.com/greenteam1/task_repo/pkg/client"
	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/repository"
)

// TaskService ...
type TaskService struct {
	TaskRepository repository.TaskRepositoryInterface
}

// NewTaskService returns created TaskService struct
func NewTaskService(tr repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{
		TaskRepository: tr,
	}
}

// GetAllTasks returns all tasks from repository
func (ts *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return ts.TaskRepository.GetAllTasks(ctx)
}

// GetTaskByID returns task with given ID
func (ts *TaskService) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	return ts.TaskRepository.GetTaskByID(ctx, id)
}

// GetTestCasesByTaskID return all test cases for task with given ID
func (ts *TaskService) GetTestCasesByTaskID(ctx context.Context, id int) ([]models.TestCase, error) {
	return ts.TaskRepository.GetTestCasesByTaskID(ctx, id)
}

// SendSolution sends solution to executioner
func (ts *TaskService) SendSolution(ctx context.Context, solutionRequest *models.SolutionSendRequest, exAddr string, taskID int) (solutionResult models.SolutionSendResponse, err error) {
	task, err := ts.TaskRepository.GetTaskByID(ctx, taskID)
	if err != nil {
		return
	}

	testCases, err := ts.TaskRepository.GetTestCasesByTaskID(ctx, task.ID)
	if err != nil {
		return
	}

	solutionResult, err = client.SendSolution(
		ctx,
		exAddr,
		solutionRequest.Language,
		solutionRequest.Code,
		task.TimeLimit,
		task.Memory,
		testCases,
	)
	return
}
