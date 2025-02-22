package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/shareed2k/goth_fiber"
	"github.com/studentkickoff/gobp/internal/api/middlewares"
	"github.com/studentkickoff/gobp/internal/api/util"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/db/repository"
	"go.uber.org/zap"
)

type AuthRouter struct {
	router   fiber.Router
	userRepo repository.User
}

func NewAPI(repo *repository.Repository, router fiber.Router) *AuthRouter {
	goth.UseProviders(
		microsoftonline.New(config.GetString("auth.msentra.client_id"), config.GetString("auth.msentra.client_secret"), config.GetString("auth.msentra.callbackURL")),
	)

	api := &AuthRouter{
		router,
		repo.NewUser(),
	}

	api.Router()

	return api
}

func (r *AuthRouter) Router() {
	auth := r.router.Group("/auth")
	auth.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	auth.Get("/callback/:provider", r.LoginCallbackHandler)
	auth.Get("/logout", r.LogoutHandler)
	auth.Get("/session", middlewares.ProtectedRoute, r.SessionHandler)
}

func (r *AuthRouter) LoginCallbackHandler(c *fiber.Ctx) error {
	// if we get logged out, we should overwrite this is with a shouldLogout = false
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		zap.L().Error("failed to complete user auth", zap.Error(err))
	}

	dbUser, err := r.userRepo.GetByUid(c.Context(), user.UserID)
	if err != nil {
		zap.L().Error("failed to search user", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	if dbUser.ID == 0 {
		if err := r.userRepo.Create(c.Context(), dbUser); err != nil {
			zap.L().Error("failed to insert user", zap.Error(err))
			return fiber.ErrInternalServerError
		}
	}

	err = util.StoreInSession("userId", dbUser.ID, c)
	if err != nil {
		zap.L().Error("failed to store user in session", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Redirect("/")
}

func (r *AuthRouter) LogoutHandler(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		zap.L().Error("failed to logout", zap.Error(err))
	}
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		zap.L().Error("failed to get session", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	if err := session.Destroy(); err != nil {
		zap.L().Error("failed to destroy", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.SendString("logout")
}

func (r *AuthRouter) SessionHandler(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
