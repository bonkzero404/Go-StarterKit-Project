package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/role/domain/dto"

	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	CreateRole(role *stores.Role) *gorm.DB

	UpdateRoleById(role *stores.Role, id string) *gorm.DB

	DeleteRoleById(role *stores.Role, id string) *gorm.DB

	GetRoleById(role *stores.Role, id string) *gorm.DB

	GetRoleList(role *stores.Role, dto *dto.RoleRequest) *gorm.DB
}
