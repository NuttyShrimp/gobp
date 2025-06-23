package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	redisstore "github.com/gofiber/storage/redis/v3"
	"github.com/shareed2k/goth_fiber"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/studentkickoff/gobp/internal/redis"
	"github.com/studentkickoff/gobp/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	*fiber.App
	Addr string
	db   database.DB
}

func NewServer() (*Server, error) {
	db, err := database.NewPSQL()

	if err != nil {
		zap.L().Error("failed to get session", zap.Error(err), zap.String("module", "database"))
		return nil, err
	}

	err = redis.New()
	if err != nil {
		zap.L().Error("Failed to connect to redis store", zap.Error(err), zap.String("module", "redis"))
		return nil, err
	}

	env := config.GetDefaultString("app.env", "development")
	var sessionStore fiber.Storage
	if env == "production" {
		sessionStore = redisstore.New(redisstore.Config{
			Host: config.GetDefaultString("redis.host", "localhost"),
			Port: config.GetDefaultInt("redis.port", 6379),
			URL:  config.GetDefaultString("redis.url", ""),
		})
	} else {
		sessionStore = postgres.New(postgres.Config{
			DB: db.Pool(),
		})
	}

	goth_fiber.SessionStore = session.New(session.Config{
		KeyLookup:      fmt.Sprintf("cookie:%s_session_id", config.GetString("app.name")),
		CookieHTTPOnly: true,
		Storage:        sessionStore,
		CookieSecure:   env == "production",
	})

	app := fiber.New()

	port := config.GetDefaultInt("server.port", 8000)
	host := config.GetDefaultString("server.host", "127.0.0.1")

	srv := &Server{
		db:   db,
		Addr: fmt.Sprintf("%s:%d", host, port),
		App:  app,
	}

	srv.RegisterRoutes()

	return srv, nil
}
