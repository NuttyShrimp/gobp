package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/studentkickoff/gobp/internal/database/dto"
	"github.com/studentkickoff/gobp/pkg/sqlc"
)

type UserRouter struct {
	router fiber.Router
	db     *sqlc.Queries
}

func NewAPI(db *sqlc.Queries, router fiber.Router) *UserRouter {
	api := &UserRouter{
		router,
		db,
	}

	return api
}

func (r *UserRouter) Router() {
	user := r.router.Group("/user")
	user.Get("/me", r.GetMeHandler)
}

func (r *UserRouter) GetMeHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	user, err := r.db.GetUser(c.Context(), userId)
	if err != nil {
		log.Error().Int32("UserID", userId).Err(err).Msg("failed to get user")
		return fiber.ErrInternalServerError
	}

	return c.JSON(dto.UserDTO(user))
}
