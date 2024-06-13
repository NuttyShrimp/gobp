package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
	"github.com/gofiber/storage/redis/v3"
	"github.com/rs/zerolog/log"
	"github.com/shareed2k/goth_fiber"
	"github.com/studentkickoff/gobp/internal/api/auth"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/sqlc"
)

type Server struct {
	*fiber.App
	Addr string
	db   *sqlc.Queries
}

func NewServer() (*Server, error) {
	db, err := database.New()

	if err != nil {
		log.Error().Str("module", "database").Err(err).Msg("")
		return nil, err
	}

	env := config.GetDefaultString("app.env", "development")
	var sessionStore fiber.Storage
	if env == "production" {
		sessionStore = redis.New()
	} else {
		sessionStore = memory.New()
	}

	goth_fiber.SessionStore = session.New(session.Config{
		KeyLookup:      fmt.Sprintf("cookie:%s_session_id", config.GetString("app.name")),
		CookieHTTPOnly: true,
		Storage:        sessionStore,
		CookieSecure:   env == "production",
	})

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
