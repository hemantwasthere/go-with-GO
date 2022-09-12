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
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursname": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
			"web dev",
			"js"
		]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json data is not valid")
	}

	// some cases where you want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %t\n", key, value, value)
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
