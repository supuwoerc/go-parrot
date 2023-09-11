package api

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/constant"
	"go-parrot/src/serializer"
	"go-parrot/src/service"
	"go-parrot/src/service/dto"
	"go-parrot/src/utils"
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
// @Success 200 {object} serializer.BasicResponse[any] "Successfully login"
// @Failure 400 {object} serializer.BasicResponse[any] "Invalid parameters"
// @Failure 500 {object} serializer.BasicResponse[any] "Internal server error"
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
// @Success 200 {object} serializer.BasicResponse[any] "Successfully add"
// @Failure 400 {object} serializer.BasicResponse[any] "Invalid parameters"
// @Failure 500 {object} serializer.BasicResponse[any] "Internal server error"
// @Router /api/public/user/add [post]
func (userApi UserApi) AddUser(ctx *gin.Context) {
	var userAddDTO dto.UserAddDTO
	err := ctx.ShouldBindJSON(&userAddDTO)
	//TODO:用户头像文件上传
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
// @Success 200 {object} serializer.BasicResponse[any] "Successfully get user info"
// @Failure 400 {object} serializer.BasicResponse[any] "Invalid parameters"
// @Failure 500 {object} serializer.BasicResponse[any] "Internal server error"
// @Router /api/public/user/{id} [get]
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
// @Param   body body dto.UserListDTO true "GET USER LIST"
// @Success 200 {object} serializer.BasicResponse[any] "Successfully get user list"
// @Failure 400 {object} serializer.BasicResponse[any] "Invalid parameters"
// @Failure 500 {object} serializer.BasicResponse[any] "Internal server error"
// @Router /api/public/user/list [post]
func (userApi UserApi) GetUserList(ctx *gin.Context) {
	paginate, paramValidErr := utils.GetPaginateParam(ctx)
	if paramValidErr != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code:    constant.InvalidParams,
			Message: paramValidErr.Error(),
		})
	} else {
		userListDTO := dto.UserListDTO{
			Paginate: paginate,
			Name:     ctx.DefaultQuery("name", ""),
		}
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

// @Tags 用户管理模块
// @Summary 更新用户信息
// @Description 更新用户信息
// @Accept  json
// @Produce  json
// @Param userUpdateDTO body dto.UserUpdateDTO true "Update User Info"
// @Success 200 {object} serializer.BasicResponse[any] "Successfully updated"
// @Failure 400 {object} serializer.BasicResponse[any] "Invalid parameters"
// @Failure 500 {object} serializer.BasicResponse[any] "Internal server error"
// @Router /api/user/update [patch]
func (userApi UserApi) UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	err := ctx.ShouldBindJSON(&userUpdateDTO)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code:    constant.InvalidParams,
			Message: err.Error(),
		})
	} else {
		err := userApi.Service.UpdateUser(&userUpdateDTO)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Data:    nil,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BasicResponse[any]{
				Code: constant.SUCCESS,
			})
		}
	}
}

// @Tags 用户管理模块
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} serializer.BasicResponse[any]
// @Failure 400 {object} serializer.BasicResponse[any]
// @Failure 500 {object} serializer.BasicResponse[any]
// @Router /api/user/delete/{id} [delete]
func (userApi UserApi) DeleteUser(ctx *gin.Context) {
	var userDeleteDTO dto.BasicIdDTO
	err := ctx.ShouldBindUri(&userDeleteDTO)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code:    constant.InvalidParams,
			Message: err.Error(),
		})
	} else {
		err := userApi.Service.DeleteUser(userDeleteDTO.ID)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Data:    nil,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BasicResponse[any]{
				Code: constant.SUCCESS,
			})
		}
	}
}
