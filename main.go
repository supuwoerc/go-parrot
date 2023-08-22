package main

import "go-parrot/src/cmd"

// @title Go-Parrot
// @version 1.0.0
// @description 鹦鹉学舌
// @contact.name Supuwoerc
// @contact.email zhangzhouou@gmail.com
// @contact.url https://github.com/supuwoerc
func main() {
	defer cmd.Clean()
	cmd.Start()
}
