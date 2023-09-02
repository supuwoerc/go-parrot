package cmd

import (
	"fmt"
	"go-parrot/src/conf"
	"go-parrot/src/global"
	"go-parrot/src/model"
	"go-parrot/src/router"
	"go-parrot/src/utils"
)

func Start() {
	var initError error = nil
	fmt.Println("cmd.Start")
	conf.InitSystemConfig()
	global.Logger = conf.InitLogger()
	database, err := conf.InitDatabase()
	if err != nil {
		initError = utils.AppendError(initError, err)
	} else {
		global.DB = database
		_ = database.AutoMigrate(&model.Request{})
	}
	redisClient, err := conf.InitRedis()
	if err != nil {
		initError = utils.AppendError(initError, err)
	} else {
		global.RedisClient = redisClient
	}
	geoIpDB, err := conf.InitGeoIpDB()
	if err != nil {
		initError = utils.AppendError(initError, err)
	} else {
		global.GeoIpDB = geoIpDB
	}
	if initError != nil {
		errorInfo := fmt.Sprintf("系统初始化失败，请检查系统配置信息：%s", initError.Error())
		global.Logger.Error(errorInfo)
		panic(errorInfo)
	}
	router.InitRouter()
}

func Clean() {
	fmt.Println("cmd.Clean")
	_ = global.GeoIpDB.Close()
}
