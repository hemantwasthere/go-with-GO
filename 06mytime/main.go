package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("Present time is", presentTime)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Present formatted time and date is", presentTime.Format("15:04:05 Monday 02-01-2006"))

	fmt.Println("----------------------------------------------------------------------")

	createdDate := time.Date(2003, time.November, 16, 04, 30, 07, 0, time.UTC)
	fmt.Println("Created date is", createdDate)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Created formatted date is", createdDate.Format("Monday 02-01-2006"))
}
