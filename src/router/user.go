package router

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/api"
)

func InitUserRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		user := api.NewUserApi()
		userGroup := public.Group("/user")
		{
			userGroup.POST("/login", user.Login)
			userGroup.POST("/add", user.AddUser)
			userGroup.GET("/:id", user.GetUserById)
		}
	})
}
