package cmd

import (
	"fmt"
	"go-parrot/src/conf"
	"go-parrot/src/global"
	"go-parrot/src/router"
)

func Start() {
	fmt.Println("cmd.Start")
	conf.InitSystemConfig()
	global.Logger = conf.InitLogger()
	router.InitRouter()
}

func Clean() {
	fmt.Println("cmd.Clean")
}
