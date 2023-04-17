package global

import (
	"go-parrot/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger //全局日志组件实例(单例)
	DB          *gorm.DB           //全局的数据库连接实例(单例)
	RedisClient *conf.RedisClient  //全局的redis连接(单例)
)
