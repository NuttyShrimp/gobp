package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"go.uber.org/zap"
)

func ProtectedRoute(c *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		zap.L().Error("failed to get session", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	if session.Fresh() {
		return c.Redirect("/login")
	}

	var userId interface{}
	if userId = session.Get("userId"); userId == nil {
		return c.Redirect("/login")
	}

	c.Locals("userId", userId)

	return c.Next()
}
