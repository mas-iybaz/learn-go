package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []string{"alpha", "beta", "charlie", "delta"}

	slice := arr[2:3]

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	// Sort an integer slice
	var sliceSample [3]int

	sliceSample = [3]int{65, 43, 2}

	random := []int{39, 17, 43, 12, 2, 78}

	fmt.Println(random)

	if !sort.IntsAreSorted(random) {
		fmt.Println("slice is not sorted yet.")
	}

	sort.Ints(random)
	sort.Ints(sliceSample[:])

	fmt.Println(random)
	fmt.Println(sliceSample)

	if sort.IntsAreSorted(random) {
		fmt.Println("slice has been sorted.")
	}
}
