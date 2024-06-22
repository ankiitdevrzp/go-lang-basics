package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://ankiit.dev:3000/learn?coursename=reactjs&paymentid=ghfd5r43"

func main() {
	fmt.Println("Welcome to URL handling in golang")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, val := range qParams {
		fmt.Println("Param is: ", val)
	}

	// always pass reference of URL not the copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "ankiit.dev",
		Path:    "/tutcss",
		RawPath: "user=ankit",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
