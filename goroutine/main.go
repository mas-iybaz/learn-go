package main

import (
	"fmt"
	"runtime"
)

func printMessage(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(i+1, message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go printMessage(25, "halo")
	printMessage(25, "hello")

	var input string
	fmt.Scanln(&input)
}
