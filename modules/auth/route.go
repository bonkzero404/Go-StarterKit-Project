package auth

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/auth/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func registerEndpointRole(name string) string {

	role, _ := driver.Casbin()

	role.AddPolicy("user", name, "", "read", "", "")

	return name
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	uriGroup := utils.SetupApiGroup() + endpointGroup
	user := app.Group(uriGroup)

	user.Post("/", middleware.RateLimiter(5, 120), handler.AuthHandler.Authentication)

	user.Get("/refresh-token", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.RefreshToken)

	user.Name(registerEndpointRole("auth.me")).Get(
		"/me",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.AuthHandler.GetProfile,
	)
}
