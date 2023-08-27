package serializer

import (
	"dario.cat/mergo"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-parrot/src/constant"
	"go-parrot/src/global"
	"net/http"
	"reflect"
)

type BasicResponse struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type DataList[T any] struct {
	List  []T `json:"list"`
	Total int `json:"total"`
}

func (res BasicResponse) IsEmpty() bool {
	return reflect.DeepEqual(res, BasicResponse{})
}

func getDefaultStatus(res BasicResponse, defaultStatus int) int {
	if 0 == res.Status {
		return defaultStatus
	}
	return res.Status
}

func HttpResponse(ctx *gin.Context, status int, res BasicResponse) {
	if res.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	if "" == res.Message {
		err := mergo.Merge(&res, BasicResponse{
			Message: constant.GetMessage(res.Code),
		})
		if err != nil {
			global.Logger.Error(fmt.Sprintf("响应数据合并失败：%s", err.Error()))
			ctx.AbortWithStatus(status)
			return
		}
	}
	ctx.AbortWithStatusJSON(status, res)
}

func Success(ctx *gin.Context, res BasicResponse) {
	HttpResponse(ctx, getDefaultStatus(res, http.StatusOK), res)
}

func Fail(ctx *gin.Context, res BasicResponse) {
	HttpResponse(ctx, getDefaultStatus(res, http.StatusBadRequest), res)
}

func ServerFail(ctx *gin.Context, res BasicResponse) {
	HttpResponse(ctx, getDefaultStatus(res, http.StatusInternalServerError), res)
}
