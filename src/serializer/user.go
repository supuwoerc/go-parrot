package serializer

import (
	"go-parrot/src/constant"
	"go-parrot/src/model"
)

type LoginSuccess struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

func BuildLoginSuccessRes(user model.User, token string) BasicResponse[LoginSuccess] {
	user.Password = ""
	return BasicResponse[LoginSuccess]{
		Code: constant.SUCCESS,
		Data: LoginSuccess{
			User:  user,
			Token: token,
		},
	}
}

func BuildUserWithoutPasswordRes(user model.User) BasicResponse[model.User] {
	user.Password = ""
	return BasicResponse[model.User]{
		Code: constant.SUCCESS,
		Data: user,
	}
}

func BuildUserListRes(list []model.User, total int64) BasicResponse[DataList[model.User]] {
	var userListWithoutPassword []model.User
	for _, val := range list {
		val.Password = ""
		userListWithoutPassword = append(userListWithoutPassword, val)
	}
	return BuildDataList[model.User](userListWithoutPassword, total)
}
