package global

import (
	"github.com/oschwald/geoip2-golang"
	"go-parrot/src/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
	GeoIpDB     *geoip2.Reader
)
