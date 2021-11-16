package services

import (
	"errors"
	"go-boilerplate-clean-arch/config"
	"go-boilerplate-clean-arch/modules/auth/domain/interfaces"
	"go-boilerplate-clean-arch/modules/auth/domain/models"
	userInterface "go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

	if errUser != nil {
		return &models.UserAuthResponse{}, errUser
	}

	match := utils.CheckPasswordHash(auth.Password, user.Password)

	if match == false {
		return &models.UserAuthResponse{}, errors.New("Invalid email or password")
	}

	claims := jwt.MapClaims{
		"id": user.ID.String(),
		// "full_name": user.FullName,
		// "email":     user.Email,
		// "phone":     user.Phone,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return &models.UserAuthResponse{}, errToken
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
		return &models.UserAuthProfileResponse{}, errUser
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
