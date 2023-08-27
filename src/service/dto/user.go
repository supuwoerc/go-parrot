package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,abc-prefix"`
	Password string `json:"password" binding:"required"`
}
