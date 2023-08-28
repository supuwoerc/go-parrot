package api

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/constant"
	"go-parrot/src/serializer"
	"go-parrot/src/service"
	"go-parrot/src/service/dto"
)

type UserApi struct {
	BasicApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BasicApi: NewBasicApi(),
		Service:  service.NewUserService(),
	}
}

// @Tags 用户管理模块
// @Summary 用户登录
// @Description 用于用户登录系统
// @Accept  json
// @Param   body body dto.UserLoginDTO true "User Login Info"
// @Success 200 {object} string "登录成功"
// @Failure 401 {object} string "登录失败"
// @Router /api/public/user/login [post]
func (userApi UserApi) Login(ctx *gin.Context) {
	var loginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&loginDTO)
	if errs != nil {
		serializer.Fail(ctx, serializer.BasicResponse{
			Code:    constant.InvalidParams,
			Message: errs.Error(),
		})
	} else {
		loginUser, token, err := userApi.Service.Login(loginDTO)
		if err != nil {
			serializer.Fail(ctx, serializer.BasicResponse{
				Code: constant.ERROR,
				Data: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BuildLoginSuccessRes(loginUser, token))
		}
	}

}
