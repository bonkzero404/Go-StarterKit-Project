package app

import (
	"go-boilerplate-clean-arch/modules/auth"
	"go-boilerplate-clean-arch/modules/user"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap(app *fiber.App) {
	user.Register(app)
	auth.Register(app)
}
