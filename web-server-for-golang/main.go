package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - lco")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:5000/get"

	res, err := http.Get(myurl)
	checkNilErr(err)
	fmt.Println("Status code is: ", res.Status)
	fmt.Println("Content length is: ", res.ContentLength)

	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	// fmt.Println("Content is: ", string(content))
	// we can do this same stuff as above using strings.Builder package -> coverting byte data to string
	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	checkNilErr(err)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())

	defer res.Body.Close()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:5000/post"

	// fake json payload or data
	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with golang",
		"price": 0,
		"platform": "learncodeonline.in"
	}
	`)
	res, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)
	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)
	fmt.Println("Content is: ", string(content))

	defer res.Body.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
