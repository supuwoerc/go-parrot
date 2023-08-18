package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IRouteRegister = func(public *gin.RouterGroup, auth *gin.RouterGroup)

var routeRegisters []IRouteRegister

// 添加路由模块
func RouteRegister(register IRouteRegister) {
	if register != nil {
		routeRegisters = append(routeRegisters, register)
	}
}

// 初始化系统模块路由
func InitBasicRouter(public *gin.RouterGroup, auth *gin.RouterGroup) {
	InitUserRoutes() //用户基础模块
	for _, val := range routeRegisters {
		val(public, auth)
	}
}

// 初始化Gin路由
func InitRouter() {
	r := gin.Default()
	publicGroup := r.Group("/api/public")
	authGroup := r.Group("/api")
	InitBasicRouter(publicGroup, authGroup)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = ":10000"
	}
	err := r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("服务启动失败：%s", err.Error()))
	}
}
