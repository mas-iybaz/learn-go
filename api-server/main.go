package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Phone struct {
	ID     string
	Name   string
	Vendor string
	Year   int
}

var data = []Phone{
	{"P001", "Redmi Note 9 Pro", "Redmi", 2020},
	{"P002", "Xiaomi Mi 11", "Xiaomi", 2021},
	{"P003", "Realme 6 Pro", "Realme", 2020},
}

func phones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var res, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func phone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" || r.Method == "GET" {
		var phoneID = r.FormValue("id")
		var res []byte
		var err error

		for _, each := range data {
			if each.ID == phoneID {
				res, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(res)
				return
			}
		}

		http.Error(w, "Phone not found!", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/phones", phones)
	http.HandleFunc("/phone", phone)

	fmt.Println("starting web server at 127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
