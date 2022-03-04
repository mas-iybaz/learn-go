package main

import (
	"fmt"
)

type Address struct {
	street, city, district, region, postCode string
	numberStreet                             int
}

func main() {
	var my_address Address

	my_address = Address{street: "Jl. Panjaitan", numberStreet: 29, city: "Madiun", district: "Jawa Timur", region: "Indonesia", postCode: "63171"}

	fmt.Println(my_address)
	fmt.Println(my_address.city)
}
