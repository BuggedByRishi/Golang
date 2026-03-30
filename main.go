package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in Go Lang")
	content := "If you are reading this it's too late"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(fliename string) {
	databyte, err := ioutil.ReadFile(fliename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
