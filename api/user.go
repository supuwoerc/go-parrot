package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录接口
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 400 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (userApi *UserApi) Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "login success!",
	})
}
