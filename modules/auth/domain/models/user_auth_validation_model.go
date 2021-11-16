package models

type UserAuthValidation struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
