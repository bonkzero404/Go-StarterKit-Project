package services

import (
	"errors"
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/domain/models"
	"go-boilerplate-clean-arch/modules/user/services/factories"
	"go-boilerplate-clean-arch/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository interfaces.UserRepositoryInterface
	ActionFactory  factories.ActionFactoryInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface, factory factories.ActionFactoryInterface) interfaces.UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
		ActionFactory:  factory,
	}
}

func (service UserService) CreateUser(user *models.UserCreateRequest) (*models.UserCreateResponse, error) {
	hashPassword, _ := utils.HashPassword(user.Password)

	userData := stores.User{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	activationCode := utils.StringWithCharset(32)

	userAvtivate := stores.UserActivation{
		Code:    activationCode,
		ActType: stores.ACTIVATION_CODE,
	}

	result, err := service.UserRepository.CreateUser(&userData, &userAvtivate)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    "User already register",
			}
		}

		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong with our server",
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": activationCode,
		},
	}

	utils.SendMail(&sendMail)

	response := models.UserCreateResponse{
		ID:       userData.ID.String(),
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		IsActive: userData.IsActive,
	}

	return &response, nil
}

func (service UserService) UserActivation(email string, code string) (*models.UserCreateResponse, error) {
	user, errUser := service.UserRepository.FindUserByEmail(email)

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	if user.IsActive {
		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "User already active",
		}
	}

	checkActivationCode, errAct := service.UserRepository.FindUserActivationCode(user.ID.String(), code)

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Activation code not found",
		}
	}

	t := time.Now()

	if checkActivationCode.ExpiredAt.Before(t) {
		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    "The activation code has expired",
		}
	}

	userNew, errUserNew := service.UserRepository.UpdateUserActivation(user.ID.String(), true)

	if errUserNew != nil {
		return &models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Cannot activated user",
		}
	}

	service.UserRepository.UpdateActivationCodeUsed(user.ID.String(), code)

	response := models.UserCreateResponse{
		ID:       userNew.ID.String(),
		FullName: userNew.FullName,
		Email:    userNew.Email,
		Phone:    userNew.Phone,
		IsActive: userNew.IsActive,
	}

	return &response, nil
}

func (service UserService) CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error) {
	user, errUser := service.UserRepository.FindUserByEmail(email)

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	if user.IsActive && actType == stores.ACTIVATION_CODE {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "User already active",
		}
	}

	_, errActFactory := service.ActionFactory.Create(actType, user)

	if errActFactory != nil {
		return nil, errActFactory
	}

	return map[string]interface{}{}, nil
}

func (service UserService) UpdatePassword(forgotPassReq *models.UserForgotPassActRequest) (map[string]interface{}, error) {
	if forgotPassReq.Password != forgotPassReq.RepeatPassword {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Password validation does not match",
		}
	}

	user, errUser := service.UserRepository.FindUserByEmail(forgotPassReq.Email)

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	checkActivationCode, errAct := service.UserRepository.FindUserActivationCode(user.ID.String(), forgotPassReq.Code)

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Activation code not found",
		}
	}

	if checkActivationCode.IsUsed {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Password reset code has been used",
		}
	}

	t := time.Now()

	if checkActivationCode.ExpiredAt.Before(t) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    "The activation code has expired",
		}
	}

	go func() {
		hashPassword, _ := utils.HashPassword(user.Password)
		service.UserRepository.UpdatePassword(user.ID.String(), hashPassword)
		service.UserRepository.UpdateActivationCodeUsed(user.ID.String(), forgotPassReq.Code)
	}()

	return map[string]interface{}{}, nil
}
