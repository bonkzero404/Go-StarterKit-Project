package role

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/role/handlers"
	"go-starterkit-project/modules/role/repositories"
	"go-starterkit-project/modules/role/services"

	"github.com/gofiber/fiber/v2"
)

/**
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	roleRepository := repositories.NewRoleRepository(driver.DB)
	roleService := services.NewRoleService(roleRepository)
	RoleHandler := handlers.NewRoleHandler(roleService)

	routesInit := ApiRoute{
		RoleHandler: *RoleHandler,
	}

	routesInit.Route(app)
}
