package handlers

import (
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/domain/models"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService interfaces.UserServiceInterface
}

func NewUserHandler(userService interfaces.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var request models.UserCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserCreateRequestValidation{
		FullName: request.FullName,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.UserService.CreateUser(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error registration", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Register user successful", response)
}

func (handler *UserHandler) UserActivation(c *fiber.Ctx) error {
	var request models.UserActivationRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserActivationRequestValidation{
		Email: request.Email,
		Code:  request.Code,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.UserService.UserActivation(request.Email, request.Code)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error activation user", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Activation user successful", response)
}

func (handler *UserHandler) ReCreateUserActivation(c *fiber.Ctx) error {
	var request models.UserReActivationRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserReActivationValidation{
		Email: request.Email,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.UserService.CreateUserActivation(request.Email, "activation_code")

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error re-create user activation", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Code activation was sent to your email", response)
}

func (handler *UserHandler) CreateActivationForgotPassword(c *fiber.Ctx) error {
	var request models.UserForgotPassRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserForgotPassValidation{
		Email: request.Email,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.UserService.CreateUserActivation(request.Email, "forgot_password")

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error create activation forgot password", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Forgot password code was sent to your email", response)
}

func (handler *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request models.UserForgotPassActRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := models.UserForgotPassActValidation{
		Email:          request.Email,
		Password:       request.Password,
		RepeatPassword: request.RepeatPassword,
		Code:           request.Code,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.UserService.UpdatePassword(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to create new password", re.StatusCode, err)
	}

	return utils.ApiCreated(c, "Successfuly to create new password", response)
}
