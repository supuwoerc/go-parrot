package main

import "go-parrot/src/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
