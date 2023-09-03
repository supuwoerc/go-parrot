package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-parrot/docs"
	"go-parrot/src/global"
	"go-parrot/src/middleware"
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
func initSystemRouter(public *gin.RouterGroup, auth *gin.RouterGroup) {
	InitBasicRoutes()          //测试路由
	InitUserRoutes()           //用户基础模块
	InitPackageManagerRoutes() //npm数据模块
	for _, val := range routeRegisters {
		val(public, auth)
	}
}

// 初始化Gin自定义验证器
func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("abc-prefix", func(fl validator.FieldLevel) bool {
			if str, ok := fl.Field().Interface().(string); ok {
				return strings.HasPrefix(str, "abc")
			}
			return false
		})
	}
}

// 初始化swagger
func initSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// 初始化Gin路由
func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	r := gin.Default()
	//跨域配置
	r.Use(middleware.Cors(), middleware.JWTMiddleware(), middleware.GeoIp())
	publicGroup := r.Group("/api/public")
	authGroup := r.Group("/api")
	initSystemRouter(publicGroup, authGroup)
	initSwagger(r)
	initValidator()
	serverPort := strings.Join([]string{":", viper.GetString("server.port")}, "")
	if serverPort == "" {
		serverPort = ":10000"
	}
	server := &http.Server{
		Addr:    serverPort,
		Handler: r,
	}
	go func() {
		global.Logger.Info(fmt.Sprintf("服务启动成功，监听端口：%s", serverPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("服务启动失败：%s", err.Error()))
			return
		}
	}()
	<-ctx.Done()
	timeoutCtx, cancelTimeoutCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeoutCtx()
	if err := server.Shutdown(timeoutCtx); err != nil {
		global.Logger.Error(fmt.Sprintf("服务强制关闭：%s", err.Error()))
		return
	}
	global.Logger.Info("服务关闭成功...")
}
