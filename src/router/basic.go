package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitBasicRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		public.POST("/ping", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	})
}
