package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Printf("Masukkan tahun kelahiran: ")
	fmt.Scanln(&input)

	var year int
	var err error

	year, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println("Tahun kelahiran:", year)
	} else {
		fmt.Println("Masukkan tahun kelahiran dengan benar")
		fmt.Println(err.Error())
	}
}
