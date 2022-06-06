package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/dto"
)

type UserServiceInterface interface {
	CreateUser(user *dto.UserCreateRequest) (*dto.UserCreateResponse, error)

	UserActivation(email string, code string) (*dto.UserCreateResponse, error)

	CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error)

	UpdatePassword(forgotPassReq *dto.UserForgotPassActRequest) (map[string]interface{}, error)
}
