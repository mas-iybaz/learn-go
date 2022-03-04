package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Printed")
	time.Sleep(time.Second * 5)
	fmt.Println("After 5 second sleep")

	// Simple Scheduler
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
		time.Sleep(time.Second * 1)
	}

	// NewTimer
	var timer = time.NewTimer(time.Second * 2)
	fmt.Println("New Timer")
	<-timer.C
	fmt.Println("Timer End")

	// AfterFunc
	var ch = make(chan bool)

	time.AfterFunc(time.Second*4, func() {
		fmt.Println("expired")
		ch <- false
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	// Ticker Scheduler
	isDone := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(time.Second * 5)
		isDone <- true
	}()

	for {
		select {
		case <-isDone:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hi", t)
		}
	}
}
