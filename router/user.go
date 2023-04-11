package router

import (
	"github.com/gin-gonic/gin"
	"go-parrot/api"
	"net/http"
)

// InitUserRouters 注册用户模块的路由
func InitUserRouters() {
	RegisterRoute(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		userApi := api.NewUserApi()
		userPublicGroup := publicRouterGroup.Group("user")
		{
			userPublicGroup.POST("/login", userApi.Login)
		}
		userAuthGroup := authRouterGroup.Group("user")
		{
			userAuthGroup.GET("", func(context *gin.Context) {
				context.AbortWithStatusJSON(http.StatusOK, gin.H{
					"message": "ok",
					"data": []map[string]interface{}{
						{"id": 1, "name": "test1"},
						{"id": 2, "name": "test2"},
					},
				})
			})
		}
	})
}
