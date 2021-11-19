package interfaces

import "go-boilerplate-clean-arch/domain/stores"

type UserForgotPassServiceFactoryInterface interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error)
}
