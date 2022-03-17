package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	studentsData := map[string]interface{}{
		"alias": map[string]string{
			"Andi":  "Alpha",
			"Budi":  "Beta",
			"Citra": "Charlie",
		},
		"total": 3,
	}

	data, _ := json.Marshal(studentsData)

	fmt.Println(string(data))
}
