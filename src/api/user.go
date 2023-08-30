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
// @Failure 500 {object} string "登录失败"
// @Router /api/public/user/login [post]
func (userApi UserApi) Login(ctx *gin.Context) {
	var loginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&loginDTO)
	if errs != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code: constant.InvalidParams,
		})
	} else {
		loginUser, token, err := userApi.Service.Login(loginDTO)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Message: err.Error(),
			})
		} else {
			//serializer.Success(ctx, serializer.BuildLoginSuccessRes(loginUser, token))
			serializer.Success(ctx, serializer.BuildLoginSuccessRes(loginUser, token))
		}
	}
}

// @Tags 用户管理模块
// @Summary 添加用户
// @Description 用于添加用户
// @Accept  json
// @Param   body body dto.UserAddDTO true "ADD USER INFO"
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string "操作失败"
// @Router /api/public/user/add [post]
func (userApi UserApi) AddUser(ctx *gin.Context) {
	var userAddDTO dto.UserAddDTO
	err := ctx.ShouldBindJSON(&userAddDTO)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code: constant.InvalidParams,
		})
	} else {
		err := userApi.Service.AddUser(&userAddDTO)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BasicResponse[dto.UserAddDTO]{
				Code: constant.SUCCESS,
				Data: userAddDTO,
			})
		}
	}
}

// @Tags 用户管理模块
// @Summary 查询用户
// @Description 用于根据ID查询用户
// @Accept  json
// @Param   body body dto.BasicIdDTO true "GET USER INFO"
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string "操作失败"
// @Router /api/public/user/add [post]
func (userApi UserApi) GetUserById(ctx *gin.Context) {
	var basicIdDTO dto.BasicIdDTO
	err := ctx.ShouldBindUri(&basicIdDTO)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code: constant.InvalidParams,
		})
	} else {
		user, err := userApi.Service.GetUserById(&basicIdDTO)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Data:    nil,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BuildUserWithoutPasswordRes(user))
		}
	}
}

// @Tags 用户管理模块
// @Summary 查询用户列表
// @Description 查询用户列表
// @Accept  json
// @Param   body body dto.BasicIdDTO true "GET USER INFO"
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string "操作失败"
// @Router /api/public/user/add [post]
func (userApi UserApi) GetUserList(ctx *gin.Context) {
	var userListDTO dto.UserListDTO
	err := ctx.ShouldBindUri(&userListDTO)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code: constant.InvalidParams,
		})
	} else {
		list, total, err := userApi.Service.GetUserList(&userListDTO)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Data:    nil,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BuildUserListRes(list, total))
		}
	}
}
