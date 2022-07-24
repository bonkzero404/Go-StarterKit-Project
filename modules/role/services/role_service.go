package services

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/role/domain/dto"
	"go-starterkit-project/modules/role/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleService struct {
	RoleRepository interfaces.RoleRepositoryInterface
}

func NewRoleService(
	roleRepository interfaces.RoleRepositoryInterface,
) interfaces.RoleServiceInterface {
	return &RoleService{
		RoleRepository: roleRepository,
	}
}

func (service RoleService) CreateRole(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error) {

	roleData := stores.Role{
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		IsActive:        true,
	}

	err := service.RoleRepository.CreateRole(&roleData).Error

	if err != nil {
		return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown", ""),
		}
	}

	roleResponse := dto.RoleResponse{
		ID:              roleData.ID.String(),
		RoleName:        roleData.RoleName,
		RoleDescription: roleData.RoleDescription,
		IsActive:        roleData.IsActive,
	}

	return &roleResponse, nil
}

func (service RoleService) GetRoleList(c *fiber.Ctx) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	res, err := service.RoleRepository.GetRoleList(&roles, c)

	if err != nil {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	for _, item := range roles {
		resp = append(resp, dto.RoleResponse{
			ID:              item.ID.String(),
			RoleName:        item.RoleName,
			RoleDescription: item.RoleDescription,
			IsActive:        item.IsActive,
		})
	}

	res.Rows = resp

	return res, nil
}
