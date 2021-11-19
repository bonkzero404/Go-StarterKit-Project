package app

import (
	"go-boilerplate-clean-arch/modules/auth"
	"go-boilerplate-clean-arch/modules/user"

	"github.com/gofiber/fiber/v2"
)

/**
This function is used to register all modules,
this registration is the last process to register
all modules
*/
func Bootstrap(app *fiber.App) {
	// Register module user
	user.RegisterModule(app)

	// Register module auth
	auth.RegisterModule(app)
}
