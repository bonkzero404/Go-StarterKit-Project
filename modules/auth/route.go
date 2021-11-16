package auth

import (
	"go-boilerplate-clean-arch/config"
	"go-boilerplate-clean-arch/modules/auth/handlers"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post("/", handler.AuthHandler.Authentication)

	user.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config("JWT_SECRET")),
	}))

	user.Get("/me", handler.AuthHandler.GetProfile)
}
