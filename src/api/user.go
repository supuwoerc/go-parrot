package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}