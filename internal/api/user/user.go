package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/studentkickoff/gobp/pkg/db/repository"
	"go.uber.org/zap"
)

type UserRouter struct {
	router   fiber.Router
	userRepo repository.User
}

func NewAPI(repo *repository.Repository, router fiber.Router) *UserRouter {
	api := &UserRouter{
		router,
		repo.NewUser(),
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

	user, err := r.userRepo.GetById(c.Context(), userId)
	if err != nil {
		zap.L().Error("failed to get user", zap.Error(err), zap.Int32("userID", userId))
		return fiber.ErrInternalServerError
	}
	if user.ID == 0 {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}
