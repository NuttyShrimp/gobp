package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/studentkickoff/gobp/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func New() (*bun.DB, error) {
	pgConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	pgConfig.ConnConfig.Host = config.GetDefaultString("db.host", "localhost")
	pgConfig.ConnConfig.Port = config.GetDefaultUint16("db.port", 5432)
	pgConfig.ConnConfig.Database = config.GetDefaultString("db.database", "postgres")
	pgConfig.ConnConfig.User = config.GetDefaultString("db.user", "postgres")
	pgConfig.ConnConfig.Password = config.GetDefaultString("db.password", "postgres")

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, err
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	// TODO: add hook to log queries in debug env

	return db, nil
}
