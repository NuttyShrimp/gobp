package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/shareed2k/goth_fiber"
)

func ProtectedRoute(c *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to get session")
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
