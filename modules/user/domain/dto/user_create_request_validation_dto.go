package dto

type UserCreateRequestValidation struct {
	FullName string `validate:"required,min=3"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric,min=10"`
}
