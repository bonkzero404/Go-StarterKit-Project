package dto

type RoleRequest struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	IsActive        bool   `json:"is_active"`
}
