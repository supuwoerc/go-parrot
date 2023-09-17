package router

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/api/npm"
)

func InitPackageManagerRoutes() {
	RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		packageManagerApi := npm.NewPackageManagerApi()
		userGroup := public.Group("/npm")
		{
			userGroup.GET("/downloads", packageManagerApi.GetDownloadsByTimeRange)
			userGroup.GET("/info", packageManagerApi.GetPackageInfo)
		}
	})
}
