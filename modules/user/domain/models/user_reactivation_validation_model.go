package models

type UserReActivationValidation struct {
	Email string `validate:"required,email"`
}
