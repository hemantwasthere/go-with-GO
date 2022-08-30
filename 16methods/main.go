package main

import "fmt"

func main() {
	fmt.Println("Welcome to study about structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("user details are: %+v\n", hitesh) // we use %+v to print the struct details in a readable format
	// fmt.Printf("The name of user is: %s\n", hitesh.Name)
	// fmt.Printf("The email of user is: %s\n", hitesh.Email)
	// fmt.Printf("The login status of user is: %v\n", hitesh.Status)
	// fmt.Printf("The age of user is: %v\n", hitesh.Age)

	hitesh.GetStatus()
	hitesh.NewMail() // this methods will not change original value of the email or struct
	fmt.Printf("The email of user is: %s\n", hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// this is method inside of the struct
func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New mail is: ", u.Email)
}
