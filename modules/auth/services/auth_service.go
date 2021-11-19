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

/**
This function is used to handle authentication
*/
func (service AuthService) Authenticate(auth *models.UserAuthRequest) (*models.UserAuthResponse, error) {
	// Get user by email
	user, errUser := service.UserRepository.FindUserByEmail(auth.Email)

	// Check if the user is not found
	// then displayan error message
	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	// Check if a query operation error occurs
	if errUser != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	// Check if the user status is not active
	if !user.IsActive {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "User is not active, please activate the user first",
		}
	}

	// Match password hashes
	match := utils.CheckPasswordHash(auth.Password, user.Password)

	// Check if it doesn't match, show an error message
	if !match {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	// Set token JWT Claims
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := jwt.MapClaims{
		"id":  user.ID.String(),
		"exp": exp,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Check if there is an error in creating the token
	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Error token",
		}
	}

	// Set response message to succeed
	response := models.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    t,
		Exp:      exp,
	}

	return &response, nil
}

/**
This function is used to authorize users and display logged in user data
*/
func (service AuthService) GetProfile(id string) (*models.UserAuthProfileResponse, error) {
	// Get user from database
	user, errUser := service.UserRepository.FindUserById(id)

	// Check if there is a query error
	if errUser != nil {
		return &models.UserAuthProfileResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	// Set response message
	response := models.UserAuthProfileResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	return &response, nil
}

/**
This function is used to refresh token
*/
func (service AuthService) RefreshToken(tokenUser *jwt.Token) (*models.UserAuthResponse, error) {
	// Get data from token then convert to string
	beforeClaims := tokenUser.Claims.(jwt.MapClaims)
	id := beforeClaims["id"].(string)

	// Get user data
	user, errUser := service.UserRepository.FindUserById(id)

	// Check if something went wrong with query
	if errUser != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	// Recreate token
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := jwt.MapClaims{
		"id":  user.ID.String(),
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return &models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Error token",
		}
	}

	// Set response message
	response := models.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    t,
		Exp:      exp,
	}

	return &response, nil
}
