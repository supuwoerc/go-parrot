package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-parrot/src/constant"
	"go-parrot/src/serializer"
	"go-parrot/src/service/dto"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// @Tags 用户管理模块
// @Summary 用户登录
// @Description 用于用户登录系统
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "登录成功"
// @Failure 401 {object} string "登录失败"
// @Router /api/public/user/login [post]
func (user User) Login(ctx *gin.Context) {
	var loginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&loginDTO)
	if errs != nil {
		fmt.Println(errs.Error())
		serializer.Fail(ctx, serializer.BasicResponse{
			Code:    constant.InvalidParams,
			Message: errs.Error(),
		})
	} else {
		serializer.Success(ctx, serializer.BasicResponse{
			Code: constant.SUCCESS,
			Data: loginDTO.Name + "欢迎你",
		})
	}

}
