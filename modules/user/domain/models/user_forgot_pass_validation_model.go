package models

type UserForgotPassValidation struct {
	Email string `validate:"required,email"`
}
