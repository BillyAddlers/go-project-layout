package main

import (
	cmd "example/cmd/example"
)

/**
* NOTE: Welcome to your Go Project!
*
* Here is your entry point, your frontline garde, the very first code to ran inside your project.
* This piece of code imports Function from `./cmd/example/main.go` for better code readability
* and to prevent cluttering inside your entry point code.
*
* Feel free to restructure your code as you see fit.
**/
func main() {
	cmd.Run()
}
