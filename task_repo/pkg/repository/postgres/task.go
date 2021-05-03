// Package postgres ...
package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// TaskRepositoryPostgresql implements models.TasksRepository
type TaskRepositoryPostgresql struct {
	db *pgxpool.Pool
}

// NewTasksRepository create new models.TasksRepository
func NewTasksRepository(db *pgxpool.Pool) *TaskRepositoryPostgresql {
	return &TaskRepositoryPostgresql{
		db: db,
	}
}

// GetAllTasks returns all tasks list from database
func (r *TaskRepositoryPostgresql) GetAllTasks(ctx context.Context) (tasks []models.Task, err error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	rows, err := r.db.Query(ctx, "SELECT id, name, description, time_limit, memory FROM tasks")
	if err != nil {
		return
	}

	for rows.Next() {
		t := models.Task{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.TimeLimit, &t.Memory)
		if err != nil {
			return
		}
		tasks = append(tasks, t)
	}
	return
}

// GetTaskByID gets task by ID
func (r *TaskRepositoryPostgresql) GetTaskByID(ctx context.Context, id int) (task models.Task, err error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	row := r.db.QueryRow(ctx, "SELECT id, name, description, time_limit, memory FROM tasks WHERE id = $1", id)

	err = row.Scan(&task.ID, &task.Name, &task.Description, &task.TimeLimit, &task.Memory)

	return
}

// GetTestCasesByTaskID gets test cases by task id
func (r *TaskRepositoryPostgresql) GetTestCasesByTaskID(ctx context.Context, taskID int) (testCases []models.TestCase, err error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	rows, err := r.db.Query(ctx, "SELECT id, task_id, test_data, answer FROM test_cases where task_id = $1", taskID)
	if err != nil {
		return
	}

	for rows.Next() {
		testCase := models.TestCase{}
		err = rows.Scan(&testCase.ID, &testCase.TaskID, &testCase.TestData, &testCase.Answer)
		if err != nil {
			return
		}
		testCases = append(testCases, testCase)
	}

	return
}
