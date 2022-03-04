package main

import (
	"fmt"
	f "fmt"
	"strings"
)

// "Call by value" function
func swapNumberValue(a, b uint8) bool {
	a, b = b, a

	return true
}

// "Call by reference" function
func swapNumberReference(a, b *uint8) bool {
	*a, *b = *b, *a

	return true
}

// Join string function -> Variadic Function
func joinString(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	var number_a, number_b uint8
	number_a, number_b = 3, 5

	f.Println("Call-by-value: swapNumber")
	f.Printf("before: number_a=%d, number_b=%d \n", number_a, number_b)
	swapNumberValue(number_a, number_b)
	f.Printf("after: number_a=%d, number_b=%d \n\n", number_a, number_b)

	f.Println("Call-by-reference: swapNumber")
	f.Printf("before: number_a=%d, number_b=%d \n", number_a, number_b)
	swapNumberReference(&number_a, &number_b)
	f.Printf("after: number_a=%d, number_b=%d \n\n", number_a, number_b)

	fmt.Println(joinString())
	fmt.Println(joinString("Muhammad", "Iqbal", "Aulia"))
}
