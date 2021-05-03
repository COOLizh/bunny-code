// Package postgres ...
package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"

	"gitlab.com/greenteam1/task_repo/pkg/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

// SolutionRepositoryPostgresql ...
type SolutionRepositoryPostgresql struct {
	db *pgxpool.Pool
}

// NewSolutionRepository ...
func NewSolutionRepository(db *pgxpool.Pool) *SolutionRepositoryPostgresql {
	return &SolutionRepositoryPostgresql{
		db: db,
	}
}

// StoreSolution stores solution into database and returns time, when it was created. The result field in database stays empty
func (r *SolutionRepositoryPostgresql) StoreSolution(ctx context.Context, uuid, lang string, userID, taskID int, sol []byte) (time.Time, error) {
	var timeStamp time.Time
	query := "INSERT INTO history_user (uuid, user_id, task_id, language_id, solution, timestamp) VALUES ($1, $2, $3, (SELECT id FROM languages WHERE name = $4), $5, current_timestamp) RETURNING timestamp;"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return timeStamp, err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, uuid, userID, taskID, lang, sol).Scan(&timeStamp)

	return timeStamp, err
}

// StoreResult updates row, containing solution, with its result
func (r *SolutionRepositoryPostgresql) StoreResult(ctx context.Context, uuid string, result models.SolutionResult) error {
	query := "UPDATE history_user SET result = $1 WHERE uuid = $2;"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, query, result, uuid)
	if err != nil {
		return err
	}

	return nil
}

// GetResultByID returns stored result by id
func (r *SolutionRepositoryPostgresql) GetResultByID(ctx context.Context, uuid string) (models.SolutionResult, error) {
	query := "SELECT result FROM history_user WHERE uuid = $1;"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return models.SolutionResult{}, err
	}
	defer conn.Release()

	var result models.SolutionResult
	err = conn.QueryRow(ctx, query, uuid).Scan(&result)
	if result.ID == "" {
		return models.SolutionResult{}, pgx.ErrNoRows
	}
	if err != nil {
		return models.SolutionResult{}, err
	}

	return result, err
}

// GetResultAndTestCaseCountByID returns stored result and amount of tests for task
func (r *SolutionRepositoryPostgresql) GetResultAndTestCaseCountByID(ctx context.Context, uuid string) (models.SolutionResult, int64, error) {
	query := "SELECT history_user.result, tc.count FROM history_user LEFT JOIN (SELECT task_id, COUNT(test_data) AS count FROM test_cases GROUP BY task_id ORDER BY task_id) tc ON history_user.task_id = tc.task_id WHERE history_user.uuid = $1;"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return models.SolutionResult{}, 0, err
	}
	defer conn.Release()

	var result models.SolutionResult
	var tcCount int64
	err = conn.QueryRow(ctx, query, uuid).Scan(&result, &tcCount)
	if result.ID == "" {
		return models.SolutionResult{}, tcCount, pgx.ErrNoRows
	}
	if err != nil {
		return models.SolutionResult{}, 0, err
	}

	return result, tcCount, err
}

// GetSolutionsHistory ...
func (r *SolutionRepositoryPostgresql) GetSolutionsHistory(ctx context.Context, userID, taskID int) ([]models.SolutionHistoryItem, error) {
	history := make([]models.SolutionHistoryItem, 0)
	query := "SELECT history_user.uuid, l.name, history_user.solution, history_user.result, history_user.timestamp FROM history_user LEFT JOIN languages l on l.id = history_user.language_id WHERE user_id = $1 AND task_id = $2 ORDER BY timestamp DESC"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return history, err
	}
	defer conn.Release()

	rows, err := r.db.Query(ctx, query, userID, taskID)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return history, err
	}

	for rows.Next() {
		item := models.SolutionHistoryItem{}
		if err := rows.Scan(
			&item.ID,
			&item.Solution.Language,
			&item.Solution.Code,
			&item.Result,
			&item.CreatedAt,
		); err != nil {
			return history, err
		}
		history = append(history, item)
	}

	return history, nil
}
