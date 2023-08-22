package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-parrot/docs"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
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
	InitBasicRoutes() //测试路由
	InitUserRoutes()  //用户基础模块
	for _, val := range routeRegisters {
		val(public, auth)
	}
}

// 初始化swagger
func InitSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// 初始化Gin路由
func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	r := gin.Default()
	publicGroup := r.Group("/api/public")
	authGroup := r.Group("/api")
	InitBasicRouter(publicGroup, authGroup)
	InitSwagger(r)
	serverPort := strings.Join([]string{":", viper.GetString("server.port")}, "")
	if serverPort == "" {
		serverPort = ":10000"
	}
	server := &http.Server{
		Addr:    serverPort,
		Handler: r,
	}
	go func() {
		fmt.Printf("服务启动成功，监听端口：%s\n", serverPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// TODO：记录日志
			fmt.Printf("服务启动失败：%s\n", err.Error())
			return
		}
	}()
	<-ctx.Done()
	timeoutCtx, cancelTimeoutCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeoutCtx()
	if err := server.Shutdown(timeoutCtx); err != nil {
		// TODO：记录日志
		fmt.Printf("服务强制关闭：%s", err.Error())
		return
	}
	// TODO：记录日志
	fmt.Println("服务关闭成功...")
}
