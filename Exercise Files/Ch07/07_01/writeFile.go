package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello from Go!"

	file, err := os.Create("C:/Users/LUCAS/Desktop/fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("All Done, File Size:", ln)

	bytes := []byte(content)

	ioutil.WriteFile("C:/Users/LUCAS/Desktop/fromString.txt", bytes, 0644)

}

func checkError(err error) {
	if err != nil {
		println(err)
		panic(err)
	}

}
