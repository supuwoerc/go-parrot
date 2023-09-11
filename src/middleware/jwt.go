package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-parrot/src/constant"
	"go-parrot/src/global"
	"go-parrot/src/model"
	"go-parrot/src/serializer"
	"go-parrot/src/service"
	"go-parrot/src/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func tokenInvalidRes(context *gin.Context) {
	serializer.Fail(context, serializer.BasicResponse[any]{
		Status: http.StatusUnauthorized,
		Code:   constant.InvalidToken,
	})
}

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader(viper.GetString("jwt.tokenKey"))
		if token == "" {
			tokenInvalidRes(context)
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil || claims.ID == 0 {
			tokenInvalidRes(context)
			return
		}
		redisKey := strings.Replace(constant.LOGIN_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(claims.ID)), -1)
		redisToken, err := global.RedisClient.Get(redisKey)
		if err != nil || redisToken != token {
			tokenInvalidRes(context)
			return
		}
		duration, err := global.RedisClient.GetExpireDuration(redisKey)
		if err != nil || duration <= 0 {
			tokenInvalidRes(context)
			return
		}
		if duration.Seconds() < (viper.GetDuration("jwt.refreshThreshold") * time.Second).Seconds() {
			refreshToken, err := utils.GenerateToken(claims.ID, claims.Name)
			if err == nil {
				err = service.SetLoginUserToken2Redis(claims.ID, refreshToken)
			}
			if err == nil {
				context.Header(viper.GetString("jwt.tokenKey"), refreshToken)
			}
		}
		context.Set(constant.LOGIN_USER_KEY, model.LoginUser{
			ID:   claims.ID,
			Name: claims.Name,
		})
	}
}
