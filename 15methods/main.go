package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")
	// no inheritance; super or parent

	ankit := User{"Ankit", "ankiit@go.dev", true, 16}
	fmt.Println(ankit)
	fmt.Printf("ankit details are: %+v\n", ankit)
	fmt.Printf("Name is %v and email is %v.", ankit.Name, ankit.Email)
	ankit.GetStatus()
	ankit.NewMail()
	fmt.Printf("Name is %v and email is %v.", ankit.Name, ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	// here it creates a copy of object; original value remains unchanged
	u.Email = "ankit@go"
	fmt.Println("Email of this user is: ", u.Email)
}
