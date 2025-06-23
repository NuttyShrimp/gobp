package api

import (
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/studentkickoff/gobp/internal/api/auth"
	"github.com/studentkickoff/gobp/internal/api/middlewares"
	"github.com/studentkickoff/gobp/internal/api/user"
	"github.com/studentkickoff/gobp/pkg/config"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoutes() {
	env := config.GetDefaultString("app.env", "development")
	s.App.Use(sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	s.App.Use(fiberzap.New(fiberzap.Config{
		Logger: zap.L(),
	}))

	api := s.App.Group("/api")

	auth.NewAPI(s.db, api)

	protectedApi := api.Use(middlewares.ProtectedRoute)

	user.NewAPI(s.db, protectedApi)

	if env != "development" {
		s.App.Static("/", "./public")
		// Fallback for SPA to handle
		s.App.Static("*", "./public/index.html")
	}

	s.App.All("/api*", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
