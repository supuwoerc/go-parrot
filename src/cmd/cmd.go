package cmd

import (
	"fmt"
	"go-parrot/src/conf"
	"go-parrot/src/router"
)

func Start() {
	fmt.Println("cmd.Start")
	conf.InitSystemConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("cmd.Clean")
}
