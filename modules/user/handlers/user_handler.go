package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/dto"
	"go-starterkit-project/modules/user/domain/interfaces"
	"go-starterkit-project/utils"

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
	var request dto.UserCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, err)
	}

	userValidation := dto.UserCreateRequestValidation{
		FullName: request.FullName,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, errors)
	}

	response, err := handler.UserService.CreateUser(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, err)
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) UserActivation(c *fiber.Ctx) error {
	var request dto.UserActivationRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, err)
	}

	userValidation := dto.UserActivationRequestValidation{
		Email: request.Email,
		Code:  request.Code,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, errors)
	}

	response, err := handler.UserService.UserActivation(request.Email, request.Code)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, err)
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) ReCreateUserActivation(c *fiber.Ctx) error {
	var request dto.UserReActivationRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, err)
	}

	userValidation := dto.UserReActivationValidation{
		Email: request.Email,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, errors)
	}

	response, err := handler.UserService.CreateUserActivation(request.Email, stores.ACTIVATION_CODE)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, err)
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) CreateActivationForgotPassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, err)
	}

	userValidation := dto.UserForgotPassValidation{
		Email: request.Email,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, errors)
	}

	response, err := handler.UserService.CreateUserActivation(request.Email, stores.FORGOT_PASSWORD)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, err)
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassActRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, err)
	}

	userValidation := dto.UserForgotPassActValidation{
		Email:          request.Email,
		Password:       request.Password,
		RepeatPassword: request.RepeatPassword,
		Code:           request.Code,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, errors)
	}

	response, err := handler.UserService.UpdatePassword(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, err)
	}

	return utils.ApiCreated(c, response)
}
