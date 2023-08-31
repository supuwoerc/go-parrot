package npm

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-parrot/src/api"
	"go-parrot/src/constant"
	"go-parrot/src/serializer"
	"go-parrot/src/service/npm"
	"go-parrot/src/utils"
	"strings"
	"time"
)

type PackageManagerApi struct {
	api.BasicApi
	Service *npm.PackageManagerService
}

func NewPackageManagerApi() PackageManagerApi {
	return PackageManagerApi{
		BasicApi: api.NewBasicApi(),
		Service:  npm.NewPackageService(),
	}
}

// @Tags NPM数据查询
// @Summary 查询NPM数最近的下载数据
// @Description 查询用户列表
// @Accept  json
// @Param   body body npm2.PackageDownloadsDTO true "GET PACKAGE DOWNLOADS"
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string "操作失败"
// @Router /api/public/npm/downloads [get]
func (packageManagerApi PackageManagerApi) GetDownloadsByTimeRange(ctx *gin.Context) {
	var err error
	var packageName string
	if packageName = ctx.Query("package"); strings.Trim(packageName, " ") == "" {
		err = utils.AppendError(err, errors.New("package_name不能为空"))
	}
	defaultStart := time.Now().AddDate(0, 0, -7).Format(time.DateOnly)
	start := ctx.DefaultQuery("start", defaultStart)
	defaultEnd := time.Now().Format(time.DateOnly)
	end := ctx.DefaultQuery("start", defaultEnd)
	if err != nil {
		serializer.Fail(ctx, serializer.BasicResponse[any]{
			Code: constant.InvalidParams,
		})
	} else {
		res, err := packageManagerApi.Service.DownloadsByTimeRange(start, end, packageName)
		if err != nil {
			serializer.ServerFail(ctx, serializer.BasicResponse[any]{
				Code:    constant.ERROR,
				Data:    nil,
				Message: err.Error(),
			})
		} else {
			serializer.Success(ctx, serializer.BasicResponse[any]{
				Data: res,
			})
		}
	}
}
