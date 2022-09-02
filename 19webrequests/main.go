package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("LCO web requests")

	res, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Type of response is %T:\n", res)

	defer res.Body.Close() // callers responsibility to close the connections

	databyes, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)
	// fmt.Println("Raw data of Body is: ", res.Body)
	content := string(databyes)
	fmt.Println("Content of Body is: ", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
