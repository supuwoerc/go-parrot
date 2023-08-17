package cmd

import (
	"fmt"
	"go-parrot/src/conf"
)

func Start() {
	fmt.Println("cmd.Start")
	conf.InitSystemConfig()
}

func Clean() {
	fmt.Println("cmd.Clean")
}
