package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/"

func main() {
	fmt.Println("Google web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}
