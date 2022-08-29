package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// simple syntax of for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for loop with range keyword
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// we can also use comma ok syntax here by replacing any value with _
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v \n", index, day)
	// }

	// rogueValue := 1
	// for rogueValue < 10 {
	// 	if rogueValue == 5 {
	// 		// break
	// 		rogueValue++ // doing rogueValue++ cuz it will not be in infinite loop
	// 		continue
	// 	}
	// 	fmt.Println("Value is:", rogueValue)
	// 	rogueValue++
	// }

	// goto statement
	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 5 {
			goto lco // using goto statement here by its label
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

	// defining goto statement label
lco:
	fmt.Println("Goto statement executes here!!!")
}
