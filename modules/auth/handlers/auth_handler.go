package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/modules/auth/domain/dto"
	"go-starterkit-project/modules/auth/domain/interfaces"
	"go-starterkit-project/utils"

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

/**
Authentication handler
*/
func (handler *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request dto.UserAuthRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := dto.UserAuthValidation{
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

/**
Get user profile
*/
func (handler *AuthHandler) GetProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	response, err := handler.AuthService.GetProfile(id)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to get user data", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Load user successful", response)
}

/**
Refresh token
*/
func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	response, err := handler.AuthService.RefreshToken(token)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to refresh token", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Refresh token successful", response)
}
