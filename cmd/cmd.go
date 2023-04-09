package cmd

import (
	"fmt"
	"go-parrot/conf"
	"go-parrot/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("程序退出完成...")
}
