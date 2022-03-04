package main

import (
	"fmt"
)

var name string

// this is init-function
func init() {
	name = "Iqbal"
}

func main() {
	// basic
	func() {
		fmt.Println("This line is printed from anonymous function")
	}()

	// store "anonymous function" to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}

	greet(name)
}
