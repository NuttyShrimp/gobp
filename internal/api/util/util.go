package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

func StoreInSession(key string, value interface{}, ctx *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(ctx)
	if err != nil {
		return err
	}

	session.Set(key, value)

	return session.Save()
}
