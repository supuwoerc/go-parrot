package main

import "go-parrot/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
