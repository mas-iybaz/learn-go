package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://api.plos.org/search?q=title:DNA"

	var u, err = url.Parse(urlString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s \n", urlString)

	fmt.Printf("Protocol: %s \n", u.Scheme)
	fmt.Printf("Host: %s \n", u.Host)
	fmt.Printf("Path: %s \n", u.Path)

	var query = u.Query()["q"][0]

	fmt.Printf("Query: %s \n", query)
}
