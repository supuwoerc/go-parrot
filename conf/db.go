package conf

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// 初始化数据库 连接
func InitDB() (*gorm.DB, error) {
	dsn := viper.GetString("db.dsn")
	isDev := viper.GetBool("mode.dev")
	loggerMode := logger.Info
	if !isDev {
		loggerMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(loggerMode),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, err
}
