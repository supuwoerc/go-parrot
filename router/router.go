package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-parrot/docs"
	"go-parrot/global"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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
	//监听系统信号量，当接收到信号量时取消上下文
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	r := gin.Default()
	publicRouterGroup := r.Group("/api/v1/public")
	authRouterGroup := r.Group("/api/v1")
	initRouterChunks()
	for _, register := range globalRouterRegisters {
		register(publicRouterGroup, authRouterGroup)
	}
	//集成swagger http://localhost:8090/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = "8848"
	}
	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: r,
	}
	go func() {
		global.Logger.Infof("服务启动，监听端口：%s", serverPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("启动服务失败：%s", err.Error())
		}
	}()
	//当ctx未取消时，程序会堵塞在此处
	<-ctx.Done()
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf("服务关闭失败：%s\n", err.Error())
	} else {
		global.Logger.Info("服务关闭成功...")
	}
}

// 按照模块来添加路由注册方法
func initRouterChunks() {
	InitUserRouters()
}
