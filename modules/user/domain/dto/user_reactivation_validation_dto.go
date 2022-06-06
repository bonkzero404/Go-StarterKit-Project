package dto

type UserReActivationValidation struct {
	Email string `validate:"required,email"`
}
