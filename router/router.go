package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRouterRegister = func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup)

// 用于注册路由的函数集合
var globalRouterRegisters []IFnRouterRegister

// RegisterRoute 添加路由注册方法
func RegisterRoute(fn IFnRouterRegister) {
	if fn != nil {
		globalRouterRegisters = append(globalRouterRegisters, fn)
	}
}

// InitRouter 注册路由模块
func InitRouter() {
	r := gin.Default()
	publicRouterGroup := r.Group("/api/v1/public")
	authRouterGroup := r.Group("/api/v1")
	initRouterChunks()
	for _, register := range globalRouterRegisters {
		register(publicRouterGroup, authRouterGroup)
	}
	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = "8848"
	}
	err := r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("服务启动失败：%s", err.Error()))
	}
}

// 按照模块来添加路由注册方法
func initRouterChunks() {
	InitUserRouters()
}
