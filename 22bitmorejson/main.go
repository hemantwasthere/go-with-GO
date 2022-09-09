package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursname"` // sets an alias for the json key
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // remove password from json data
	Tags     []string `json:"tags,omitempty"` // remove all the null fields
}

func main() {
	fmt.Println("Welcome to handle json data in golang")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// packet this data as JSON data
	// finalJson, err := json.Marshal(lcoCourses)
	// finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t") // MarshalIndent(interface name, prefix, indent based on what)
	checkNilErr(err)
	fmt.Printf("Final json is: %s\n", finalJson)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
