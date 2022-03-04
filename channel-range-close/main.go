package main

import (
	"fmt"
	"runtime"
)

/*
	note: Channel Direction
	ch chan int		: can send and receive
	ch chan<- int	: can send only
	ch <-chan int	: can receive only
*/

//  ch chan<- string : only send data
func sendMessage(ch chan<- string, iteration int) {
	for i := 0; i < iteration; i++ {
		ch <- fmt.Sprintf("data %d main", i+1)
	}
	close(ch)
}

// ch <-chan string : only receive data
func logMessage(ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {
	runtime.GOMAXPROCS(3)

	var message = make(chan string)

	go sendMessage(message, 25)
	logMessage(message)
}
