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
// @Summary 获取指定时间范围内的下载数据
// @Description 根据指定的时间范围和包名获取下载数据
// @Param package query string true "包名"
// @Param start query string false "开始日期 (默认为7天前)" Format(yyyy-mm-dd)
// @Param end query string false "结束日期 (默认为今天)" Format(yyyy-mm-dd)
// @Success 200 {object} BasicResponse{data=DownloadData} "成功返回数据"
// @Failure 400 {object} BasicResponse{code=int,message=string} "请求参数错误"
// @Failure 500 {object} BasicResponse{code=int,message=string} "服务器错误"
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
