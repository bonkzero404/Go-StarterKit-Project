package dto

type RoleRequestValidation struct {
	RoleName string `validate:"required"`
	IsActive bool   `validate:"required"`
}
