package main

import "fmt"

// jwtToken := 'lsdjflksd23kklsk32' // no var style dont work outside of method
const LoginToken string = "lsdjflksd23kklsk32" // this is public variable cuz of putting first letter as uppercase

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n\n", username)

	var isLoogedIn bool = false
	fmt.Println(isLoogedIn)
	fmt.Printf("Variable is of type: %T \n\n", isLoogedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n\n", smallValue)

	// var smallFloat float32 = 255.39293823923892 --> float32 will take upto 5 decimals only
	var smallFloat float64 = 255.39293823923892
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n\n", smallFloat)

	// default values and some aliases
	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n\n", anotherVar)

	// implicit way of declaring variables --> lexer automatically detects the type on the basis of what value you are assigning to it
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n\n", website)

	// no var style --> inside any method we can declare variables like this
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n\n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n\n", LoginToken)

}
