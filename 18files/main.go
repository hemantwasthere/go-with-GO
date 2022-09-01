package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./createdFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./createdFile.txt")
}

func readFile(filePath string) {
	databyte, err := ioutil.ReadFile(filePath)
	checkNilErr(err)
	// fmt.Println("Row data is: \n", databyte) // row data in byte
	fmt.Println("Text data inside the file is: \n", string(databyte)) // converted row data into string
}

// checking error functions
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
