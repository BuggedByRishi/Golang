package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn"

func main() {
	fmt.Println("Welcome to handeling URL's in Golang")
	fmt.Println(myurl)

	// Parsing the URL

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
}
