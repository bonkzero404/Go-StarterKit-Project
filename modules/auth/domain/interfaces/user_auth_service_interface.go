package interfaces

import "go-boilerplate-clean-arch/modules/auth/domain/models"

type UserAuthServiceInterface interface {
	Authenticate(auth *models.UserAuthRequest) (*models.UserAuthResponse, error)
	GetProfile(id string) (*models.UserAuthProfileResponse, error)
}
