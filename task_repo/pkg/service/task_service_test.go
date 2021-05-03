package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/models/mocks"
)

func TestGetTaskByID(t *testing.T) {
	mockPGTaskRepo := mocks.PGTaskRepoMock{}
	taskID := 1
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockTask := mocks.SetMockTask()

		mockPGTaskRepo.On(
			"GetTaskByID",
			mock.AnythingOfType("int"),
		).Return(mockTask, nil).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		task, err := taskServiceMock.GetTaskByID(ctx, taskID)
		assert.NoError(t, err)
		assert.Exactly(t, mockTask, task)
		assert.EqualValues(t, mockTask.ID, task.ID)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGTaskRepo.On(
			"GetTaskByID",
			mock.AnythingOfType("int"),
		).Return(models.Task{}, errors.New("Unexpexted Error")).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		task, err := taskServiceMock.GetTaskByID(ctx, taskID)

		assert.Error(t, err)
		assert.Exactly(t, models.Task{}, task)
	})
}

func TestGetTestCasesByTaskID(t *testing.T) {
	mockPGTaskRepo := mocks.PGTaskRepoMock{}
	taskID := 1
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockTestCase := mocks.SetMockTestCase()
		mockListTestCase := make([]models.TestCase, 0)
		mockListTestCase = append(mockListTestCase, mockTestCase)

		mockPGTaskRepo.On(
			"GetTestCasesByTaskID",
			mock.AnythingOfType("int"),
		).Return(mockListTestCase, nil).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		testCaseList, err := taskServiceMock.GetTestCasesByTaskID(ctx, taskID)
		assert.NoError(t, err)
		assert.Len(t, testCaseList, len(testCaseList))
	})

	t.Run("error-failed", func(t *testing.T) {
		testCasesEmpthy := make([]models.TestCase, 0)
		mockPGTaskRepo.On(
			"GetTestCasesByTaskID",
			mock.AnythingOfType("int"),
		).Return(testCasesEmpthy, errors.New("Unexpexted Error")).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		testCaseList, err := taskServiceMock.GetTestCasesByTaskID(ctx, taskID)

		assert.Error(t, err)
		assert.Len(t, testCaseList, 0)
	})
}

func TestGetAllTask(t *testing.T) {
	mockPGTaskRepo := mocks.PGTaskRepoMock{}
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockTask := mocks.SetMockTask()
		mockListTask := make([]models.Task, 0)
		mockListTask = append(mockListTask, mockTask)

		mockPGTaskRepo.On(
			"GetAllTasks",
		).Return(mockListTask, nil).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		list, err := taskServiceMock.GetAllTasks(ctx)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListTask))
	})

	t.Run("error-failed", func(t *testing.T) {
		var tasks []models.Task
		mockPGTaskRepo.On(
			"GetAllTasks",
		).Return(tasks, errors.New("Unexpexted Error")).Once()

		taskServiceMock := NewTaskService(&mockPGTaskRepo)

		list, err := taskServiceMock.GetAllTasks(ctx)

		assert.Error(t, err)
		assert.Len(t, list, 0)
	})
}
