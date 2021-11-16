package interfaces

import (
	"go-boilerplate-clean-arch/domain/stores"
)

type UserRepositoryInterface interface {
	CreateUser(user *stores.User) (*stores.User, error)
}
