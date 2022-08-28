package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Lets see if else in golang")

	var result string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the login count: ")
	input, _ := reader.ReadString('\n')

	loginCount, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// fmt.Printf("Type of loginCount is %T\n", loginCount)
	// fmt.Printf("Value of loginCount is: %v\n", loginCount)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		if loginCount < 10 {
			result = "Regular User"
		} else if loginCount > 10 {
			result = "Admin User"
		} else {
			result = "Moderate User"
		}
		fmt.Println(result)
	}

	// if else with short statement
	// if 9%2 == 0 {
	// 	fmt.Println("Number is even")
	// } else {
	// 	fmt.Println("Number is odd")
	// }

	// // initializing variable in if else and further checking it
	// if num := 3; num < 10 {
	// 	fmt.Println("Number is less than 10")
	// } else {
	// 	fmt.Println("Number is not less than 10")
	// }

}
