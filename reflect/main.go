package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 369

	reflectNumber := reflect.ValueOf(number)

	fmt.Println("Type:", reflectNumber.Type())
	fmt.Println("Value:", reflectNumber.Int())
}
