package interfaces

import (
	"go-boilerplate-clean-arch/modules/auth/domain/data_models"

	"github.com/golang-jwt/jwt/v4"
)

type UserAuthServiceInterface interface {
	Authenticate(auth *data_models.UserAuthRequest) (*data_models.UserAuthResponse, error)

	GetProfile(id string) (*data_models.UserAuthProfileResponse, error)

	RefreshToken(token *jwt.Token) (*data_models.UserAuthResponse, error)
}
