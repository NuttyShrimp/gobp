package db

import (
	"context"

	"github.com/studentkickoff/gobp/pkg/db/sqlc"
)

// DB represents a database connection
type DB interface {
	Queries() *sqlc.Queries
	WithRollback(ctx context.Context, fn func(q *sqlc.Queries) error) error
}
