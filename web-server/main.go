package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func greet(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":  "Muhammad Iqbal Aulia",
		"Greet": "Welcome",
	}

	var t, err = template.ParseFiles("index.htm")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Homepage!")
	})

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/greet", greet)

	fmt.Println("starting web server at 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
