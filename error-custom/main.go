package main

import (
	"errors"
	"fmt"
	"strings"
)

func isValidate_trim(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Field required!")
	}

	return true, nil
}

func main() {
	var name string

	fmt.Printf("Input your name: ")
	fmt.Scanln(&name)

	if valid_name, err := isValidate_trim(name); valid_name {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println(err.Error())
	}
}
