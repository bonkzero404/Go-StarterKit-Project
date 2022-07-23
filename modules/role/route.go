package role

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/role/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	RoleHandler handlers.RoleHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	role := app.Group(utils.SetupApiGroup() + endpointGroup)

	role.Post("/", middleware.RateLimiter(5, 30), handler.RoleHandler.CreateRole)

	role.Get("/", middleware.RateLimiter(5, 30), handler.RoleHandler.GetRoleList)

}
