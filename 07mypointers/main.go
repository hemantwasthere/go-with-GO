package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("The default value of ptr is", ptr)
	// fmt.Printf("The type of ptr is %T", ptr)

	myNumber := 23
	ptr := &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	// *ptr *= 2 // same as above
	fmt.Println("Value of modified pointer is", *ptr)
}
