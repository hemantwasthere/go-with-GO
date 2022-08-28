package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to study about maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages are:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "PY") // this will delete the key value pair from the map
	fmt.Println("List of all languages after deletion are:", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	// in case you dont care about then you can use comma ok syntax
	for _, value := range languages {
		fmt.Printf("Value is %v\n", value)
	}

}
