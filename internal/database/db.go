package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/studentkickoff/gobp/pkg/sqlc"
)

type DB interface {
	WithRollback(ctx context.Context, fn func(q *sqlc.Queries) error) error
	Pool() *pgxpool.Pool
	Queries() *sqlc.Queries
}
