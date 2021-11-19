package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/models"
)

type UserServiceInterface interface {
	CreateUser(user *models.UserCreateRequest) (*models.UserCreateResponse, error)

	UserActivation(email string, code string) (*models.UserCreateResponse, error)

	CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error)

	UpdatePassword(forgotPassReq *models.UserForgotPassActRequest) (map[string]interface{}, error)
}
