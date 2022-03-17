package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Phone struct {
	ID     string
	Name   string
	Vendor string
	Year   int
}

var baseUrl = "http://127.0.0.1:8080"

func fetchPhones() ([]Phone, error) {
	var err error
	var client = &http.Client{}
	var data []Phone

	req, err := http.NewRequest("GET", baseUrl+"/phones", nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchPhone(ID string) (Phone, error) {
	var err error
	var client = &http.Client{}
	var data Phone

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	req, err := http.NewRequest("POST", baseUrl+"/phone", payload)
	if err != nil {
		return data, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var phones, err = fetchPhones()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	var phone, e = fetchPhone("P003")
	if e != nil {
		fmt.Println("Error!", e.Error())
		return
	}

	for _, each := range phones {
		fmt.Println(each)
	}

	fmt.Printf("ID: %s \t Name: %s \t Vendor: %s", phone.ID, phone.Name, phone.Vendor)
}
