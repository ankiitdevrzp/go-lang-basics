package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This will go in a file - github.com"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close() // defer is recommended
	readFile("./mygofile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Text data inside file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
