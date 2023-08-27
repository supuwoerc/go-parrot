package serializer

import (
	"go-parrot/src/constant"
	"go-parrot/src/model"
)

type LoginSuccess struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func BuildLoginSuccessRes(user model.User, token string) BasicResponse {
	return BasicResponse{
		Code: constant.SUCCESS,
		Data: LoginSuccess{
			Name:  user.Name,
			Token: token,
		},
	}
}
