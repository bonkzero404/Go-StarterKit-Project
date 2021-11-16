package user

import (
	"go-boilerplate-clean-arch/modules/user/handlers"
	"go-boilerplate-clean-arch/modules/user/repositories"
	"go-boilerplate-clean-arch/modules/user/services"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
