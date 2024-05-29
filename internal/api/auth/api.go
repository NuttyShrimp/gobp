package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type API struct {
	router fiber.Router
}

func NewAPI(db *bun.DB, router fiber.Router) *API {
	api := &API{
		router: router,
	}

	return api
}

func (a *API) Router() {
	// auth := a.router.Group("/auth")
	// auth.Get("/auth/login/:provider")
}
