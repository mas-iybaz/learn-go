package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height float32
	Width  float32
}

// 'omitempty' -> ignoring empty fields
// `json:"-"`  -> alwayys ignore a field
type Smartphone struct {
	Name       string
	Vendor     string `json:"provider,omitempty"`
	Dimensions Dimensions
}

func main() {
	phone := &Smartphone{
		Name:   "Redmi Note 9 Pro",
		Vendor: "Redmi",
		Dimensions: Dimensions{
			Height: 20,
			Width:  9,
		},
	}

	data, _ := json.Marshal(phone)

	fmt.Println(string(data))

	phone_2 := &Smartphone{
		Name:   "Realme 6 Pro",
		Vendor: "Realme",
		Dimensions: Dimensions{
			Height: 21,
			Width:  9.3,
		},
	}

	data_2, _ := json.Marshal([]*Smartphone{phone, phone_2})

	fmt.Println(string(data_2))
}
