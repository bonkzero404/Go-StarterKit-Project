package handlers

import (
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/modules/auth/domain/interfaces"
	"go-boilerplate-clean-arch/modules/auth/domain/models"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		EmailValid:    request.Email,
		PasswordValid: request.Password,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.AuthService.Authenticate(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Invalid authentication", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Authentication successful", response)
}

func (handler *AuthHandler) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	response, err := handler.AuthService.GetProfile(id)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to get user data", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Load user successful", response)
}
