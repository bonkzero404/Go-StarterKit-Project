package interfaces

import "go-boilerplate-clean-arch/modules/user/domain/models"

type UserServiceInterface interface {
	CreateUser(user *models.UserCreateRequest) (*models.UserCreateResponse, error)
}
