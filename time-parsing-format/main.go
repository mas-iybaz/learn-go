package main

import (
	"fmt"
	"time"
)

func main() {
	var timeNow = time.Now()

	fmt.Printf("time is %v \n", timeNow)
	fmt.Printf("day: %v \n", timeNow.Weekday())
	fmt.Printf("date: %v \n", timeNow.Day())
	fmt.Printf("month: %v \n", timeNow.Month())
	fmt.Printf("year: %v \n", timeNow.Year())

	// ===========================================

	var timestampFormat, value string
	var date time.Time

	timestampFormat = "2006-01-02 15:04:05 MST"
	value = "2022-03-01 08:04:00 WIB"
	// timestampFormat = "Jan 2, 2006 at 3:04pm"
	// value = "Feb 28, 2022 at 9:39pm"

	date, _ = time.Parse(timestampFormat, value)

	fmt.Println(value, "\t->", date.String())

	// format date to string
	var dateFormat = date.Format("Monday, 02 January 2006")

	fmt.Println(dateFormat)
}
