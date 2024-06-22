package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance; super or parent

	ankit := User{"Ankit", "ankiit@go.dev", true, 16}
	fmt.Println(ankit)
	fmt.Printf("ankit details are: %+v\n", ankit)
	fmt.Printf("Name is %v and email is %v.", ankit.Name, ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
