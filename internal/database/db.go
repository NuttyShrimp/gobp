package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/sqlc"
)

func New() (*sqlc.Queries, *pgxpool.Pool, error) {
	pgConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, nil, err
	}
	pgConfig.ConnConfig.Host = config.GetDefaultString("db.host", "localhost")
	pgConfig.ConnConfig.Port = config.GetDefaultUint16("db.port", 5432)
	pgConfig.ConnConfig.Database = config.GetDefaultString("db.database", "gobp")
	pgConfig.ConnConfig.User = config.GetDefaultString("db.user", "postgres")
	pgConfig.ConnConfig.Password = config.GetDefaultString("db.password", "postgres")

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, nil, err
	}

	if err := pool.Ping(context.TODO()); err != nil {
		return nil, nil, err
	}

	queries := sqlc.New(pool)

	return queries, pool, nil
}
