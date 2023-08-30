package serializer

import (
	"go-parrot/src/constant"
	"go-parrot/src/model"
)

type LoginSuccess struct {
	Name  string `json:"name"`
	Token string `json:"token"`
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
