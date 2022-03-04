package main

import (
	"fmt"
	"runtime"
)

func printMessage(msg chan string) {
	fmt.Println(<-msg)
}

func main() {
	runtime.GOMAXPROCS(3)

	var guests = make(chan string)

	var greet = func(name string) {
		fmt.Println("Hello,", name)

		guests <- name
	}

	go greet("Sutrisno")
	go greet("Siti Maimunah")
	go greet("Muhammad Iqbal Aulia")
	go greet("Haidar Aziz Habibulloh")

	var guest1 = <-guests
	fmt.Println(guest1)

	var guest2 = <-guests
	fmt.Println(guest2)

	var guest3 = <-guests
	fmt.Println(guest3)

	var guest4 = <-guests
	fmt.Println(guest4)

	// Pass as a parameter
	var students = make(chan string)
	var lists = []string{"Alpha", "Beta", "Charlie", "Delta"}
	for _, each := range lists {
		go func(name string) {
			var data = fmt.Sprintf("Hi %s", name)
			students <- data
		}(each)
	}

	for i := 0; i < len(lists); i++ {
		printMessage(students)
	}
}
