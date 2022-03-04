package main

import (
	"fmt"
)

func main() {
	// normal for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for act like while
	i := 0
	for i <= 7 {
		fmt.Println(i)
		i += 2
	}

	data := map[string]string{
		"name":   "Iqbal",
		"city":   "Madiun",
		"region": "Indonesia",
	}

	for key, value := range data {
		fmt.Println(key, value)
	}
}
