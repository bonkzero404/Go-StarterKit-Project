package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
)

type UserRepositoryInterface interface {
	CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error)

	FindUserByEmail(email string) (*stores.User, error)

	FindUserById(id string) (*stores.User, error)

	FindUserActivationCode(userId string, code string) (*stores.UserActivation, error)

	UpdateUserActivation(id string, stat bool) (*stores.User, error)

	ReCreateUserActivation(userActivate *stores.UserActivation) (*stores.UserActivation, error)
}
