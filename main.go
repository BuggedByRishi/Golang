package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Hrushikesh",
		"surname": "Kakulte",
	}

	fmt.Println(person["name"], person["surname"])
}
