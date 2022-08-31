package main

import "fmt"

func main() {
	fmt.Println("Lets see defer in golang")

	// it will be deffered in the reverse order like (LIFO order)
	defer fmt.Println("World") // by putting defer here, it will be executed at the end of the program
	defer fmt.Println("One")   // by putting defer here, it will be executed at the end of the program
	defer fmt.Println("Two")   // by putting defer here, it will be executed at the end of the program
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
