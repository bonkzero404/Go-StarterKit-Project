package handlers

import (
	"go-boilerplate-clean-arch/modules/auth/domain/interfaces"
	"go-boilerplate-clean-arch/modules/auth/domain/models"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService interfaces.UserAuthServiceInterface
}

func NewAuthHandler(authService interfaces.UserAuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (handler *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request models.UserAuthRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserAuthValidation{
		Email:    request.Email,
		Password: request.Password,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.AuthService.Authenticate(&request)

	if err != nil {
		return utils.ApiUnprocessableEntity(c, "Invalid authentication", err)
	}

	return utils.ApiOk(c, "Authentication successful", response)
}
