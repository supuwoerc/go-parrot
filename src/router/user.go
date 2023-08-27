package router

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/api"
)

func InitUserRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		userGroup := public.Group("/user")
		{
			user := api.NewUserApi()
			userGroup.POST("/login", user.Login)
		}

	})
}
