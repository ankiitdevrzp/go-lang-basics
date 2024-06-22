package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is starting...")

	error := http.ListenAndServe(":4000", r)
	fmt.Println("Server connect: ", error.Error())

	log.Fatal(error)

	fmt.Println("Listening on port 4000...")
}
