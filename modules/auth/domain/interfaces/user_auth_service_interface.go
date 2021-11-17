package interfaces

import (
	"go-boilerplate-clean-arch/modules/auth/domain/models"

	"github.com/golang-jwt/jwt/v4"
)

type UserAuthServiceInterface interface {
	Authenticate(auth *models.UserAuthRequest) (*models.UserAuthResponse, error)

	GetProfile(id string) (*models.UserAuthProfileResponse, error)

	RefreshToken(token *jwt.Token) (*models.UserAuthResponse, error)
}
