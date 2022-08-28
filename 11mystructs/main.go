package main

import "fmt"

func main() {
	fmt.Println("Welcome to study about structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("user details are: %+v", hitesh) // we use %+v to print the struct details in a readable format
	// fmt.Printf("The name of user is: %s\n", hitesh.Name)
	// fmt.Printf("The email of user is: %s\n", hitesh.Email)
	// fmt.Printf("The login status of user is: %v\n", hitesh.Status)
	// fmt.Printf("The age of user is: %v\n", hitesh.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
