package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"` // Custom field name
	Age      int
}

func main() {
	var jsonString = `{"Name": "Muhammad Iqbal", "Age": 22}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("fullname:", data.Fullname)
	fmt.Println("age:", data.Age)

	// ==========> Decode JSON -> map[string]interface{} & interface{}
	var data_1 map[string]interface{}
	json.Unmarshal(jsonData, &data_1)

	fmt.Println("fullname:", data_1["Name"])
	fmt.Println("age:", data_1["Age"])

	var data_2 interface{}
	json.Unmarshal(jsonData, &data_2)

	var decodedData = data_2.(map[string]interface{})

	fmt.Println("fullname:", decodedData["Name"])
	fmt.Println("age:", decodedData["Age"])

	// =============> Decode JSON array to Object array
	var jsonStringNew = `[
		{"Name": "Muhammad Iqbal", "Age": 22},
		{"Name": "Haidar Aziz", "Age": 8}
	]`

	var data_3 []User

	var e = json.Unmarshal([]byte(jsonStringNew), &data_3)
	if e != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1 name:", data_3[0].Fullname)
	fmt.Println("user 2 name:", data_3[1].Fullname)

	// =============> Encode Object to JSON
	var obj = []User{
		{"Alpha", 9},
		{"Beta", 9},
	}

	var jsonObj, er = json.Marshal(obj)

	if er != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonStr = string(jsonObj)
	fmt.Println(jsonStr)
}
