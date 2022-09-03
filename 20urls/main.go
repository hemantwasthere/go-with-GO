package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dkgalkdf"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println("URL is: ", myurl)

	// parsing the url
	result, err := url.Parse(myurl)
	checkNilErr(err)
	// fmt.Println("Result is:", result)
	// fmt.Println("Result Scheme is:", result.Scheme)
	// fmt.Println("Result Host is:", result.Host)
	// fmt.Println("Result Path is:", result.Path)
	// fmt.Println("Result RawQuery is:", result.RawQuery)
	// fmt.Println("Result Port is:", result.Port())

	qparams := result.Query()
	fmt.Printf("Type of Query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Query params is: ", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url is: ", anotherUrl)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
