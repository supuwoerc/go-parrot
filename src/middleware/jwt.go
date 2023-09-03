package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//TODO:解析token设置Gin.Context用户信息内容
	}
}
