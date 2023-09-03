package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// 拼接错误
func AppendError(prev, next error) error {
	if prev == nil {
		return next
	}
	return fmt.Errorf("%v,%w", prev, next)
}

// 从请求中获取分页参数
func GetPaginateParam(ctx *gin.Context) (int, int, error) {
	var paramValidErr error
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "1")
	if pageTrimStr := strings.Trim(pageStr, " "); pageTrimStr == "" {
		pageStr = "1"
	}
	if pageSizeTrimStr := strings.Trim(pageSizeStr, " "); pageSizeTrimStr == "" {
		pageSizeStr = "10"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		paramValidErr = AppendError(paramValidErr, err)
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		paramValidErr = AppendError(paramValidErr, err)
	}
	return page, pageSize, err
}
