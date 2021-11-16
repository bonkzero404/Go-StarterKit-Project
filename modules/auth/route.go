package auth

import (
	"go-boilerplate-clean-arch/infrastructure/middleware"
	"go-boilerplate-clean-arch/modules/auth/handlers"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post("/", handler.AuthHandler.Authentication)

	user.Get("/me", middleware.Authenticate(), handler.AuthHandler.GetProfile)
}
