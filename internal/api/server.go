package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/studentkickoff/gobp/internal/api/auth"
	"github.com/studentkickoff/gobp/internal/database"
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

	viper.SetDefault("server.port", 8000)
	viper.SetDefault("server.host", "127.0.0.1")
	port := viper.GetInt("server.port")
	host := viper.GetString("server.host")

	srv := &Server{
		db:   db,
		Addr: fmt.Sprintf("%s:%d", host, port),
		App:  app,
	}
	return srv, nil
}
