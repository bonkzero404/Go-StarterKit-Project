package data_models

type UserCreateRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
