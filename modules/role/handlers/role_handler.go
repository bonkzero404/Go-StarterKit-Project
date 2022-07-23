package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/modules/role/domain/dto"
	"go-starterkit-project/modules/role/domain/interfaces"
	"go-starterkit-project/modules/role/domain/validation"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	RoleService interfaces.RoleServiceInterface
}

func NewRoleHandler(roleService interfaces.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{
		RoleService: roleService,
	}
}

func (handler *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	roleValidation := validation.RoleRequestValidation{
		RoleName: request.RoleName,
	}

	errors := utils.ValidateStruct(roleValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.RoleService.CreateRole(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error registration", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Register user successful", response)
}

func (handler *RoleHandler) GetRoleList(c *fiber.Ctx) error {
	response, err := handler.RoleService.GetRoleList(c)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error get roles", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Get role successful", response)
}
