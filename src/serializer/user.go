package serializer

import (
	"go-parrot/src/constant"
	"go-parrot/src/model"
	"gorm.io/gorm"
)

type LoginSuccess struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserWithoutPassword struct {
	gorm.Model
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func BuildLoginSuccessRes(user model.User, token string) BasicResponse[LoginSuccess] {
	return BasicResponse[LoginSuccess]{
		Code: constant.SUCCESS,
		Data: LoginSuccess{
			Name:  user.Name,
			Token: token,
		},
	}
}

func BuildUserWithoutPasswordRes(user model.User) BasicResponse[UserWithoutPassword] {
	return BasicResponse[UserWithoutPassword]{
		Code: constant.SUCCESS,
		Data: UserWithoutPassword{
			Model:    user.Model,
			Name:     user.Name,
			RealName: user.RealName,
			Avatar:   user.Avatar,
			Phone:    user.Phone,
			Email:    user.Phone,
		},
	}
}

func BuildUserListRes(list []model.User, total int64) BasicResponse[DataList[model.User]] {
	return BuildDataList[model.User](list, total)
}
