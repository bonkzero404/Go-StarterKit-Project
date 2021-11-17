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

	user.Post("/", middleware.RateLimiter(5, 120), handler.AuthHandler.Authentication)

	user.Get("/me", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.GetProfile)

	user.Get("/refresh-token", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.RefreshToken)
}
