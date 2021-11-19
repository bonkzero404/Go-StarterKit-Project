package services

import (
	"errors"
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/domain/models"
	"go-boilerplate-clean-arch/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
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
		Code: activationCode,
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

	response := models.UserCreateResponse{
		ID:       userNew.ID.String(),
		FullName: userNew.FullName,
		Email:    userNew.Email,
		Phone:    userNew.Phone,
		IsActive: userNew.IsActive,
	}

	return &response, nil

}
