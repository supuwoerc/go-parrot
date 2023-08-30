package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserAddDTO struct {
	ID       uint   `form:"id" json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	RealName string `form:"real_name" json:"real_name"`
	Avatar   string `form:"avatar" json:"avatar"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}

type UserListDTO struct {
	Paginate
	Name string `form:"name" json:"name"`
}
