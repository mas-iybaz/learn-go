package main

import (
	"fmt"
)

type Student struct {
	name        string
	height      float64
	age         uint8
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = Student{
		name:        "Iqbal",
		height:      5.7,
		age:         22,
		isGraduated: true,
		hobbies:     []string{"coding", "gaming"},
	}

	// %b	-> numeric to binary (string)
	fmt.Printf("%d is %b \n", data.age, data.age)

	// %c	-> numeric to unicode (string)
	fmt.Printf("%d is %c \n", 1235, 1235)

	// %e	-> numeric to scientific notation (string)
	fmt.Printf("%.2f is %e \n", data.height, data.height)

	// %o	-> numeric to octaf (string)
	fmt.Printf("%d is %o \n", data.age, data.age)

	// %t	-> boolean
	fmt.Printf("isGraduated is %t \n", data.isGraduated)

	// %T	-> data type
	fmt.Printf("isGraduated is %T \n", data.isGraduated)

	// %+v	-> struct
	fmt.Printf("data: %+v \n", data)

	// %#v	-> struct like as when declaration
	fmt.Printf("data: %#v \n", data)

	// %%
	fmt.Printf("data: %%d \n")
}
