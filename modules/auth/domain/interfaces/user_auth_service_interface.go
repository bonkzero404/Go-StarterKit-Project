package interfaces

import (
	"go-starterkit-project/modules/auth/domain/dto"

	"github.com/golang-jwt/jwt/v4"
)

type UserAuthServiceInterface interface {
	Authenticate(auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error)

	GetProfile(id string) (*dto.UserAuthProfileResponse, error)

	RefreshToken(token *jwt.Token) (*dto.UserAuthResponse, error)
}
