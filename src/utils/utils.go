package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-parrot/src/service/dto"
)

// 拼接错误
func AppendError(prev, next error) error {
	if prev == nil {
		return next
	}
	return fmt.Errorf("%v,%w", prev, next)
}

// 从请求中获取分页参数
func GetPaginateParam(ctx *gin.Context) (dto.Paginate, error) {
	var params dto.Paginate
	if err := ctx.ShouldBindQuery(&params); err != nil {
		return params, err
	}
	return params, nil
}
