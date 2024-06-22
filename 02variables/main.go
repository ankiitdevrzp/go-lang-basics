package main

import (
	"fmt"
)

// jwtToken := 300000 -> this can't be declared outside methods
// var jwtToken = 300000 -> this is allowed

const LoginToken string = "abcdefgh" // L in caps make the variable PUBLIC and can't used in any file inside the same package

func main() {
	var username string = "ankit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45554645644564546454
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable id of type: %T \n", anotherVariable)

	// implicit type
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numberOfuser := 300000.0
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable id of type: %T \n", LoginToken)
}
