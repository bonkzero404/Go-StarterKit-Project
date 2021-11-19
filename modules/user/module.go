package user

import (
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/handlers"
	"go-boilerplate-clean-arch/modules/user/repositories"
	"go-boilerplate-clean-arch/modules/user/services"
	"go-boilerplate-clean-arch/modules/user/services/factories"

	"github.com/gofiber/fiber/v2"
)

func registerActivationFactory(userRepository interfaces.UserRepositoryInterface) factories.ActionFactoryInterface {
	actFactory := factories.NewUserActivationServiceFactory(userRepository)
	forgotPassFactory := factories.NewUserForgotPassServiceFactory(userRepository)

	return factories.NewActionFactory(actFactory, forgotPassFactory)
}

func RegisterModule(app *fiber.App) {

	userRepository := repositories.NewUserRepository()

	c := registerActivationFactory(userRepository)

	userService := services.NewUserService(userRepository, c)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
