package dto

type UserForgotPassActRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Code           string `json:"code"`
}
