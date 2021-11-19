package auth

import (
	"go-boilerplate-clean-arch/modules/auth/handlers"
	"go-boilerplate-clean-arch/modules/auth/services"
	"go-boilerplate-clean-arch/modules/user/repositories"

	"github.com/gofiber/fiber/v2"
)

/**
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {
	userRepository := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	routesInit := ApiRoute{
		AuthHandler: *authHandler,
	}

	routesInit.Route(app)
}
