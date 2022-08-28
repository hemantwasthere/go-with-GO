package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on array")

	var fruitList [4]string
	// We can also directly initialize an array like this
	// var fruitList = [4]string{"Apple", "Banana", "Grape", "Tomato"}

	fruitList[0] = "Apple"
	// fruitList[1] = "Banana"
	fruitList[2] = "Grape"
	fruitList[3] = "Tomato"

	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("Length of Fruit List is:", len(fruitList))

	var vegList = [5]string{"Carrot", "Potato", "Onion"}
	fmt.Println("Veg List is:", vegList)
	fmt.Println("Length of Veg List is:", len(vegList))
}
