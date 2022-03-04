package main

import (
	"fmt"
	"runtime"
)

func getMin(numbers []int, ch chan int) {
	var min = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	ch <- min
}

func getAvg(numbers []int, ch chan float64) {
	var sum = 0

	for _, num := range numbers {
		sum += num
	}

	var avg = float64(sum) / float64(len(numbers))

	ch <- avg
}

func main() {
	runtime.GOMAXPROCS(3)

	var numbers = []int{7, 6, 4, 9, 3, 5, 6, 2, 8, 1}
	fmt.Println("numbers:", numbers)

	var ch1 = make(chan int)
	go getMin(numbers, ch1)

	var ch2 = make(chan float64)
	go getAvg(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case min := <-ch1:
			fmt.Println("Min number:", min)
		case avg := <-ch2:
			fmt.Println("Average:", avg)
		}
	}
}
