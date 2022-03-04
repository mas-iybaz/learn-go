package main

import "fmt"

func main() {
	// datatype: number
	// var number_1 int
	// var number_2 int8
	// var number_3 int16
	// var number_4 int32
	// var number_5 int64
	// var number_6 uint
	// var number_7 uint8
	// var number_8 uint16
	// var number_9 uint32
	// var number_10 uint64
	// var number_11 rune // Synonym of int32
	// var number_12 byte // Synonym of uint8
	// var number_13 uintptr
	// var number_14 float32
	// var number_15 float64
	// var number_16 complex64
	// var number_17 complex128

	// datatype: boolean
	// var isActive bool

	// datatype: string
	// var name string

	var name string
	var age uint8
	var isLearn bool

	name = "Muhammad Iqbal Aulia"
	age = 22
	isLearn = true

	if isLearn {
		fmt.Printf("Hello, my name is %v, I'm %d years old", name, age)
	}
}
