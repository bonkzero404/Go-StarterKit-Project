package validation

type RoleRequestValidation struct {
	RoleName string `validate:"required"`
}
