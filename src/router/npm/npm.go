package npm

import (
	"github.com/gin-gonic/gin"
	"go-parrot/src/api/npm"
	"go-parrot/src/router"
)

func InitPackageManagerRoutes() {
	router.RouteRegister(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		packageManagerApi := npm.NewPackageManagerApi()
		userGroup := public.Group("/npm")
		{
			userGroup.GET("/downloads", packageManagerApi.GetDownloadsByTimeRange)
		}
	})
}
