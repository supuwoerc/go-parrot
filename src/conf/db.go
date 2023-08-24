package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go-parrot/src/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDatabase() *gorm.DB {
	dsn := viper.GetString("mysql.dsn")
	logMode := logger.Info
	if !viper.GetBool("mode.dev") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		errorInfo := fmt.Sprintf("数据库连接失败,请检查dsn设置：%s", err.Error())
		global.Logger.Error(errorInfo)
		panic(errorInfo)
	}
	link, err := db.DB()
	if err != nil {
		errorInfo := fmt.Sprintf("获取DB连接失败：%s", err.Error())
		global.Logger.Error(errorInfo)
		panic(errorInfo)
	}
	link.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConn"))
	link.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConn"))
	link.SetConnMaxLifetime(time.Hour)
	return db
}
