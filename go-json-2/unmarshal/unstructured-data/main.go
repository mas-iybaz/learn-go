package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var responseJson = `{
		"status": 200,
		"message": "OK",
		"token": "J765YT90908032022",
		"phone": {
			"name": "Redmi Note 9 Pro",
			"vendor": "Redmi",
			"dimensions": {
				"height": 20,
				"width": 9
			}
		}
	}`

	var response map[string]interface{}

	json.Unmarshal([]byte(responseJson), &response)

	phone := response["phone"].(map[string]interface{})
	dimensions := phone["dimensions"].(map[string]interface{})

	for key, value := range dimensions {
		fmt.Println(key, value.(float64))
	}
}
