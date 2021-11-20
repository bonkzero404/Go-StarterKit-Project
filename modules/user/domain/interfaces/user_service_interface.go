package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/data_models"
)

type UserServiceInterface interface {
	CreateUser(user *data_models.UserCreateRequest) (*data_models.UserCreateResponse, error)

	UserActivation(email string, code string) (*data_models.UserCreateResponse, error)

	CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error)

	UpdatePassword(forgotPassReq *data_models.UserForgotPassActRequest) (map[string]interface{}, error)
}
