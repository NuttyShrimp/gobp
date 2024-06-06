package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/rs/zerolog/log"
	"github.com/shareed2k/goth_fiber"
	"github.com/spf13/viper"
	"github.com/studentkickoff/gobp/internal/api/util"
	"github.com/studentkickoff/gobp/internal/database"
	"github.com/uptrace/bun"
)

type API struct {
	router fiber.Router
	db     *bun.DB
}

func NewAPI(db *bun.DB, router fiber.Router) *API {
	goth.UseProviders(
		microsoftonline.New(viper.GetString("auth.msentra.client_id"), viper.GetString("auth.msentra.client_secret"), viper.GetString("auth.msentra.callbackURL")),
	)

	api := &API{
		router,
		db,
	}

	return api
}

func (a *API) Router() {
	auth := a.router.Group("/auth")
	auth.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	auth.Get("/callback/:provider", a.LoginCallbackHandler)
	auth.Get("/logout", a.LogoutHandler)
}

func (a *API) LoginCallbackHandler(c *fiber.Ctx) error {
	// if we get logged out, we should overwrite this is with a shouldLogout = false
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to complete user auth")
	}

	dbUser := &database.User{}
	err = a.db.NewSelect().Model(user).Where("uid = ?", user.UserID).Scan(c.Context(), &dbUser)
	if err != nil {
		log.Error().Err(err).Msg("failed to search user")
		return fiber.ErrInternalServerError
	}

	if dbUser.Id == 0 {
		dbUser = &database.User{Name: user.Name, Uid: user.UserID, Email: user.Email}
		_, err := a.db.NewInsert().Model(dbUser).Exec(c.Context())
		if err != nil {
			log.Fatal().Err(err).Msg("failed to insert user")
			return fiber.ErrInternalServerError
		}
	}

	util.StoreInSession("userId", dbUser.Id, c)

	return nil
}

func (a *API) LogoutHandler(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Error().Err(err).Msg("failed to logout")
	}

	return c.SendString("logout")
}
