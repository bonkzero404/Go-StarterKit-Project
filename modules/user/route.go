package user

import (
	"go-boilerplate-clean-arch/modules/user/handlers"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	UserHandler handlers.UserHandler
}

func (controller *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post("/register", controller.UserHandler.RegisterUser)
}
