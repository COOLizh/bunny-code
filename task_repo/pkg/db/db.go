// Package db implements mock database functions
package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// Connect function gets a database connection
func Connect(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	if err := conn.Conn().Ping(ctx); err != nil {
		return nil, err
	}
	log.Println("Ping database success")

	return pool, nil
}
