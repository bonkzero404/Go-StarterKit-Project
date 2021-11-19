package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
)

type UserActivationServiceFactoryInterface interface {
	CreateUserActivation(user *stores.User) (*stores.UserActivation, error)
}
