package repositories

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/role/domain/dto"
	"go-starterkit-project/modules/role/domain/interfaces"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) interfaces.RoleRepositoryInterface {
	return &RoleRepository{
		DB: db,
	}
}

func (repository RoleRepository) CreateRole(role *stores.Role) *gorm.DB {
	return repository.DB.Create(&role)
}

func (repository RoleRepository) UpdateRoleById(role *stores.Role, id string) *gorm.DB {
	return repository.DB.Save(&role)
}

func (repository RoleRepository) DeleteRoleById(role *stores.Role, id string) *gorm.DB {
	return repository.DB.First(&role, "email = ?", id)
}

func (repository RoleRepository) GetRoleById(role *stores.Role, id string) *gorm.DB {
	return repository.DB.First(&role, "id = ?", id)
}

func (repository RoleRepository) GetRoleList(role *stores.Role, dto *dto.RoleRequest) *gorm.DB {
	return repository.DB.Where(&stores.Role{
		RoleName: dto.RoleName,
	}).Find(&role)
}
