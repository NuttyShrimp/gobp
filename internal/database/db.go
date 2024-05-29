package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func New() (*bun.DB, error) {
	viper.BindEnv("db.host", "DB_HOST")
	viper.SetDefault("db.host", "localhost")
	viper.BindEnv("db.port", "DB_PORT")
	viper.SetDefault("db.port", 5432)
	viper.BindEnv("db.user", "DB_USER")
	viper.SetDefault("db.user", "postgres")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.SetDefault("db.password", "postgres")
	viper.BindEnv("db.database", "DB_DATABASE")
	viper.SetDefault("db.database", "postgres")

	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.ConnConfig.Host = viper.GetString("db.host")
	config.ConnConfig.Port = viper.GetUint16("db.port")
	config.ConnConfig.Database = viper.GetString("db.database")
	config.ConnConfig.User = viper.GetString("db.user")
	config.ConnConfig.Password = viper.GetString("db.password")

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
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
