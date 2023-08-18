package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		public.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "login success",
			})
		})
	})
}
