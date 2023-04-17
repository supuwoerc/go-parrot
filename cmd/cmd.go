package cmd

import (
	"fmt"
	"go-parrot/conf"
	"go-parrot/global"
	"go-parrot/router"
	"go-parrot/utils"
)

func Start() {
	//初始化项目中产生的全部错误
	var initErrors error
	//初始化配置文件
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()
	//初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErrors = utils.AppendError(initErrors, err)
	}
	redisClient, err := conf.InitRedis()
	global.RedisClient = redisClient
	if err != nil {
		initErrors = utils.AppendError(initErrors, err)
	}
	//收集初始化过程中产生的全部错误
	if initErrors != nil {
		if global.Logger != nil {
			global.Logger.Error(initErrors.Error())
		}
		panic(initErrors.Error())
	}
	//初始化gin-router
	router.InitRouter()
}

func Clean() {
	fmt.Println("程序退出清理动作完成...")
}
