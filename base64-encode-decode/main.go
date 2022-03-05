package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var name = "Muhammad Iqbal Aulia"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)

	var decodedString = string(decodedByte)

	fmt.Println(decodedString)

	// =================>>> URL encoding
	var url = "https://emiya.id/"

	var encodedUrl = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedUrl)

	var decodedByteUrl, _ = base64.URLEncoding.DecodeString(encodedUrl)
	var decodedUrl = string(decodedByteUrl)

	fmt.Println(decodedUrl)
}
