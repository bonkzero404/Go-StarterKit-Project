package interfaces

import (
	"go-starterkit-project/modules/role/domain/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleServiceInterface interface {
	CreateRole(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleList(c *fiber.Ctx) (*utils.Pagination, error)

	// UpdateRole(role *dto.RoleRequest) (*dto.RoleResponse, error)

	// DeleteRoleById(id string) (map[string]interface{}, error)
}
