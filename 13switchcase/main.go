package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to learn about switch case in golang")

	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	diceNumber := rand.Intn(6) + 1   // generating random number between 1 and 6
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dive value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
		// fallthrough // fallthrough keyword is used to execute the next case
	case 3:
		fmt.Println("You can move 3 spots")
		// fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again!!")
	default:
		fmt.Println("What was that!!!")
	}
}
