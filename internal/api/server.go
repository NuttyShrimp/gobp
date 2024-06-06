package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/studentkickoff/gobp/internal/api/auth"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/uptrace/bun"
)

type Server struct {
	*fiber.App
	Addr string
	db   *bun.DB
}

func NewServer() (*Server, error) {
	db, err := database.New()
	if err != nil {
		log.Error().Str("module", "database").Err(err).Msg("")
		return nil, err
	}

	app := fiber.New()
	api := app.Group("/api")

	authAPI := auth.NewAPI(db, api)
	authAPI.Router()

	port := config.GetDefaultInt("server.port", 8000)
	host := config.GetDefaultString("server.host", "127.0.0.1")

	srv := &Server{
		db:   db,
		Addr: fmt.Sprintf("%s:%d", host, port),
		App:  app,
	}
	return srv, nil
}
