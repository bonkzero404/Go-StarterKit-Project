package user

import (
	"go-boilerplate-clean-arch/app/middleware"
	"go-boilerplate-clean-arch/modules/user/handlers"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	UserHandler handlers.UserHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post("/register", middleware.RateLimiter(5, 30), handler.UserHandler.RegisterUser)

	user.Post("/activation", middleware.RateLimiter(5, 30), handler.UserHandler.UserActivation)

	user.Post("/activation/re-send", middleware.RateLimiter(5, 30), handler.UserHandler.ReCreateUserActivation)

	user.Post("/request-forgot-password", middleware.RateLimiter(5, 30), handler.UserHandler.CreateActivationForgotPassword)

	user.Post("/forgot-password", middleware.RateLimiter(5, 30), handler.UserHandler.UpdatePassword)
}
