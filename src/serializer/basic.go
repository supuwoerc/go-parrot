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

type BasicResponse[T any] struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type DataList[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

func (res BasicResponse[any]) IsEmpty() bool {
	return reflect.DeepEqual(res, BasicResponse[any]{})
}

func getDefaultStatus[T any](res BasicResponse[T], defaultStatus int) int {
	if 0 == res.Status {
		return defaultStatus
	}
	return res.Status
}

func HttpResponse[T any](ctx *gin.Context, status int, res BasicResponse[T]) {
	if res.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	if "" == res.Message {
		err := mergo.Merge(&res, BasicResponse[any]{
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

func BuildDataList[T any](list []T, total int64) BasicResponse[DataList[T]] {
	return BasicResponse[DataList[T]]{
		Code: constant.SUCCESS,
		Data: DataList[T]{
			List:  list,
			Total: total,
		},
	}
}

func Success[T any](ctx *gin.Context, res BasicResponse[T]) {
	HttpResponse[T](ctx, getDefaultStatus(res, http.StatusOK), res)
}

func Fail[T any](ctx *gin.Context, res BasicResponse[T]) {
	HttpResponse[T](ctx, getDefaultStatus(res, http.StatusBadRequest), res)
}

func ServerFail[T any](ctx *gin.Context, res BasicResponse[T]) {
	HttpResponse[T](ctx, getDefaultStatus(res, http.StatusInternalServerError), res)
}
