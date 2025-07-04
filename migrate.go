package main

import (
	"embed"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/studentkickoff/gobp/pkg/config"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	// setup database
	db, err := database.NewPSQL()
	if err != nil {
		panic(err)
	}
	conn := stdlib.OpenDBFromPool(db.Pool())

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(conn, "db/migrations"); err != nil {
		panic(err)
	}

	// run app
}
