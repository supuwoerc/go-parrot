package conf

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDatabase() (*gorm.DB, error) {
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
		return nil, err
	}
	link, err := db.DB()
	if err != nil {
		return nil, err
	}
	link.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConn"))
	link.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConn"))
	link.SetConnMaxLifetime(time.Hour)
	return db, err
}
