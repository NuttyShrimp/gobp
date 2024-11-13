package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/studentkickoff/gobp/internal/database/dto"
	"github.com/studentkickoff/gobp/pkg/sqlc"
	"go.uber.org/zap"
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

	api.Router()

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
		zap.L().Error("failed to get user", zap.Error(err), zap.Int32("userID", userId))
		return fiber.ErrInternalServerError
	}

	return c.JSON(dto.UserDTO(user))
}
