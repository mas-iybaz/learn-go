package main

import "fmt"

func main() {
	// this is a CONSTANT, not a variable
	const EARTH_GRAVITY = 10

	// using var
	var name string
	var city, district string
	name = "Muhammad Iqbal Aulia"
	city, district = "Madiun", "East Java"

	// short variable declaration
	age := 22

	fmt.Println(EARTH_GRAVITY)
	fmt.Println(name)
	fmt.Println(city)
	fmt.Println(district)
	fmt.Println(age)
}
