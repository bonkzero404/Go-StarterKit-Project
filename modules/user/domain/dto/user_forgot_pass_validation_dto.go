package dto

type UserForgotPassValidation struct {
	Email string `validate:"required,email"`
}
