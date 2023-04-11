package main

import "go-parrot/cmd"

// @title GO-Web开发记录
// @version 1.0.0
// @description 使用golang开发后台管理系统
func main() {
	defer cmd.Clean()
	cmd.Start()
}
