package models

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Gender          int    `json:"gender" binding:"oneof=0 1 2"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
