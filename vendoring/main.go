package main

import (
	"fmt"

	termtable "github.com/scylladb/termtables"
)

type smartphone struct {
	Name   string
	Vendor string
	Year   int
}

func main() {
	var phones []smartphone = []smartphone{
		{"Redmi Note 9 Pro", "Redmi", 2020},
		{"Samsung Galaxy S22 Plus", "Samsung", 2021},
		{"iPhone 15 Plus", "Apple", 2022},
	}

	table := termtable.CreateTable()

	table.AddHeaders("Name", "Vendor", "Year")

	for _, phone := range phones {
		table.AddRow(phone.Name, phone.Vendor, phone.Year)
	}

	fmt.Println(table.Render())
}
