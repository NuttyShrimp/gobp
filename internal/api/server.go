package api

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/storage/redis/v3"
	"github.com/shareed2k/goth_fiber"
	"github.com/studentkickoff/gobp/internal/api/auth"
	"github.com/studentkickoff/gobp/internal/api/middlewares"
	"github.com/studentkickoff/gobp/internal/api/user"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/sqlc"
	"go.uber.org/zap"
)

type Server struct {
	*fiber.App
	Addr string
	db   *sqlc.Queries
}

func NewServer() (*Server, error) {
	db, pool, err := database.New()

	if err != nil {
		zap.L().Error("failed to get session", zap.Error(err), zap.String("module", "database"))
		return nil, err
	}

	env := config.GetDefaultString("app.env", "development")
	var sessionStore fiber.Storage
	if env == "production" {
		sessionStore = redis.New()
	} else {
		sessionStore = postgres.New(postgres.Config{
			DB: pool,
		})
	}

	goth_fiber.SessionStore = session.New(session.Config{
		KeyLookup:      fmt.Sprintf("cookie:%s_session_id", config.GetString("app.name")),
		CookieHTTPOnly: true,
		Storage:        sessionStore,
		CookieSecure:   env == "production",
	})

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetString("app.dsn"),
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	app := fiber.New()
	app.Use(sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: zap.L(),
	}))

	api := app.Group("/api")

	authAPI := auth.NewAPI(db, api)
	authAPI.Router()

	protectedApi := api.Use(middlewares.ProtectedRoute)

	userAPI := user.NewAPI(db, protectedApi)
	userAPI.Router()

	if env != "development" {
		app.Static("/", "./public")
		// Fallback for SPA to handle
		app.Static("*", "./public/index.html")
	}

	app.All("/api*", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := config.GetDefaultInt("server.port", 8000)
	host := config.GetDefaultString("server.host", "127.0.0.1")

	srv := &Server{
		db:   db,
		Addr: fmt.Sprintf("%s:%d", host, port),
		App:  app,
	}
	return srv, nil
}
