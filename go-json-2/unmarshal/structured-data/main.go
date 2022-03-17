package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
}

type Student struct {
	Name string
	Age  int
}

type Dimensions struct {
	Height float32
	Width  float32
}

type Smartphone struct {
	Name       string
	Vendor     string
	Dimensions Dimensions
}

func main() {
	// Decoding JSON into structs
	var userJson = `{"username": "iqbal"}`
	var user User

	json.Unmarshal([]byte(userJson), &user)

	fmt.Println("username", user.Username)

	// Array JSON
	var studentsJson = `[
		{"name": "Alpha", "age": 17},
		{"name": "Beta", "age": 16},
		{"name": "Charlie", "age": 18}
	]`
	var students []Student

	json.Unmarshal([]byte(studentsJson), &students)

	for _, student := range students {
		fmt.Println("name:", student.Name)
		fmt.Println("age:", student.Age)
		fmt.Println()
	}

	// Nested JSON object
	var phoneJson = `{
		"name": "Redmi Note 9 Pro",
		"vendor": "Redmi",
		"dimensions": {
			"height": 20.5,
			"width": 9.7
		}
	}`
	var phone Smartphone

	json.Unmarshal([]byte(phoneJson), &phone)

	fmt.Println("smarphone name:", phone.Name)
	fmt.Println("smarphone vendor:", phone.Vendor)
	fmt.Println("smarphone height:", phone.Dimensions.Height)
	fmt.Println("smarphone width:", phone.Dimensions.Width)
}
