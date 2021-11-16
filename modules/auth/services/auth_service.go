package services

import (
	"errors"
	"go-boilerplate-clean-arch/config"
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/modules/auth/domain/interfaces"
	"go-boilerplate-clean-arch/modules/auth/domain/models"
	userInterface "go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository userInterface.UserRepositoryInterface
}

func NewAuthService(userRepository userInterface.UserRepositoryInterface) interfaces.UserAuthServiceInterface {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service AuthService) Authenticate(auth *models.UserAuthRequest) (*models.UserAuthResponse, error) {
	user, errUser := service.UserRepository.FindUserByEmail(auth.Email)

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	if errUser != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	match := utils.CheckPasswordHash(auth.Password, user.Password)

	if match == false {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	claims := jwt.MapClaims{
		"id":  user.ID.String(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Error token",
		}
	}

	response := models.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    t,
	}

	return &response, nil
}

func (service AuthService) GetProfile(id string) (*models.UserAuthProfileResponse, error) {
	user, errUser := service.UserRepository.FindUserById(id)

	if errUser != nil {
		return &models.UserAuthProfileResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	response := models.UserAuthProfileResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	return &response, nil
}
