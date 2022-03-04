package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
	defer fmt.Println("action ended")
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("Receive data", data)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout! no activity under 5 seconds")
			break loop
		}
	}
}

func main() {
	runtime.GOMAXPROCS(3)
	rand.Seed(time.Now().Unix())

	var message = make(chan int)

	go sendData(message)
	retreiveData(message)

	// Exit
	os.Exit(1)

	fmt.Println("This statement is not excecuted!")
}
