package main

import (
	"fmt"
	"learnmodule/robo"
)

func main() {
	var name string

	fmt.Printf("Please input your name: ")
	fmt.Scanln(&name)

	robo.Greet(name)
}
