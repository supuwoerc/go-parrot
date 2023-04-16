package cmd

import (
	"fmt"
	"go-parrot/conf"
	"go-parrot/global"
	"go-parrot/router"
)

func Start() {
	//初始化配置文件
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()
	//初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		fmt.Println(err.Error())
	}
	//初始化gin-router
	router.InitRouter()
}

func Clean() {
	fmt.Println("程序退出清理动作完成...")
}
