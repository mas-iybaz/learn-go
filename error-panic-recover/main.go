package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(name string) (bool, error) {
	defer fmt.Println("Validation ended")

	if strings.TrimSpace(name) == "" {
		return false, errors.New("Name cannot be empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Program running perfectly")
	}
}

func main() {
	defer catch()

	var name string

	fmt.Printf("Input your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		panic(err.Error())
	}

	fmt.Println("Bye")
	fmt.Println("Oh.. One more")

	// Recover in IIFE
	superheroes := []string{"superman", "batman", "wonder woman", " ", "aquaman"}

	for _, each := range superheroes {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic Error on looping", each, "| msg:", r)
				} else {
					fmt.Println("App running normally")
				}
			}()

			if valid, err := validate(each); valid {
				fmt.Println("Hello", each)
			} else {
				panic(err.Error())
			}

			panic("some error happen")
		}()
	}
}
