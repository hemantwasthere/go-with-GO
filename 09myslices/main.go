package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to learn about slices which is more often used in golang")
	fmt.Println("------------------------------------------------------------------")

	var fruitList = []string{"Apple", "Banana", "Grape", "Tomato"}
	fmt.Printf("Fruit List is: %v of type %T\n", fruitList, fruitList)

	fruitList = append(fruitList, "Orange")
	fmt.Println("Appended Fruit List is:", fruitList)

	// fruitList = append(fruitList[1:]) // slicing the fruitList from 1st index to end
	fruitList = append(fruitList[1:3]) // slicing the fruitList from 1st index to 3rd where 3rd is excluded
	fmt.Println("Sliced Fruit List is:", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	highScores[3] = 400
	// highScores[4] = 500  this will throw an error as we have only 4 elements limit

	highScores = append(highScores, 500, 700, 600) // but we can append highScores like this

	fmt.Println("High Scores are:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores are:", highScores)
	fmt.Println("Is highScores sorted or not?", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"Docker", "Kubernetes", "Puppet", "Terraform", "AWS", "Azure", "GCP"}
	fmt.Println("All Courses are:", courses)
	index := 2
	// first we sliced the slice from 0 to index and then we sliced from index+1 to end
	courses = append(courses[:index], courses[index+1:]...) // this will remove the value at index 2
	fmt.Println("Remaining Courses after deletion are:", courses)

}
