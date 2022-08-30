package main

import "fmt"

func main() {
	fmt.Println("Lets see functions in golang")
	greeter()

	result := adder(2, 3)
	fmt.Println("Result coming from adder is :", result)

	// using variadic function, also we can use comma ok syntax here
	proResInt, proResString := proAdder(2, 3, 4)
	// fmt.Println("Result coming from proAdder is :", proAdder(1, 2, 3, 4))
	fmt.Println("Int result coming from proAdder is :", proResInt)
	fmt.Println("String result coming from proAdder is :", proResString)
}

func proAdder(value ...int) (int, string) { // specifying here that it will return int and string
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "Hi this is proAdder"
}

func adder(val1 int, val2 int) int { // this is function signature, the type of returned value
	return val1 + val2
}

func greeter() {
	fmt.Println("Namaste from go lang")
}
