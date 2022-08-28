package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// fmt.Println("Thanks for rating, ", strings.TrimSpace(input)+"huh")

	// strings.TrimSpace(input) is used to remove the newline character from the input and trim the string
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // converting string to float with 64 bit precision

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Added 1 to your rating, ", numRating+1)
	}

}
