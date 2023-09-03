package router

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/api"
)

func InitUserRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		user := api.NewUserApi()
		userGroup := public.Group("/user")
		userAuthGroup := auth.Group("/user")
		{
			userGroup.POST("/login", user.Login)
			userGroup.POST("/add", user.AddUser)
			userGroup.GET("/:id", user.GetUserById)
		}
		{
			userAuthGroup.GET("/list", user.GetUserList)
			userAuthGroup.PATCH("/update", user.UpdateUser)
		}
	})
}
