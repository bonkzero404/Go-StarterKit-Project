package user

import (
	"go-boilerplate-clean-arch/modules/user/aggregates"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/handlers"
	"go-boilerplate-clean-arch/modules/user/repositories"
	"go-boilerplate-clean-arch/modules/user/services"
	"go-boilerplate-clean-arch/modules/user/services/factories"

	"github.com/gofiber/fiber/v2"
)

/**
Service factory registration
*/
func registerActivationFactory(userActivationRepository interfaces.UserActivationRepositoryInterface) factories.ActionFactoryInterface {
	actFactory := factories.NewUserActivationServiceFactory(userActivationRepository)
	forgotPassFactory := factories.NewUserForgotPassServiceFactory(userActivationRepository)

	return factories.NewActionFactory(actFactory, forgotPassFactory)
}

/**
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	userRepository := repositories.NewUserRepository()
	userActivationRepository := repositories.NewUserActivationRepository()
	aggregateRepository := aggregates.NewRepositoryAggregate(userRepository, userActivationRepository)

	userActivationFactory := registerActivationFactory(userActivationRepository)

	userService := services.NewUserService(userRepository, userActivationRepository, aggregateRepository, userActivationFactory)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
