package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)

	// messages := make(chan int, 1)

	// go func() {
	// 	for {
	// 		i := <-messages
	// 		fmt.Println("Receive data", i)
	// 	}
	// }()

	// for i := 0; i < 9; i++ {
	// 	fmt.Println("Send data", i)
	// 	messages <- i
	// }

	drinks := make(chan string, 1)

	go func() {
		for {
			i := <-drinks
			fmt.Println(i)
		}
	}()

	for _, each := range []string{"Milku", "Floridina", "Teh Pucuk Harum", "Isoplus", "Fruit Tea", "Kopikap"} {
		fmt.Println("Load data:", each)
		drinks <- each
	}

	var input string
	fmt.Scanln(&input)
}
