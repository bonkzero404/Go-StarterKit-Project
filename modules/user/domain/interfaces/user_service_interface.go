package interfaces

import "go-boilerplate-clean-arch/modules/user/domain/models"

type UserServiceInterface interface {
	CreateUser(user *models.UserCreateRequest) (*models.UserCreateResponse, error)

	UserActivation(email string, code string) (*models.UserCreateResponse, error)

	ReCreateUserActivation(email string) (map[string]interface{}, error)
}
