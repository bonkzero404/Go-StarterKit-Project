package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
)

type UserRepositoryInterface interface {
	CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error)

	FindUserByEmail(email string) (*stores.User, error)

	FindUserById(id string) (*stores.User, error)
}
