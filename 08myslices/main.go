package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // range is non inclusive, it will take 1 to 2 here
	// fruitList = append(fruitList[:3]) // it will take 0 to 2 here
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 245
	highScores[1] = 259
	highScores[2] = 247
	highScores[3] = 248
	// highScores[4] = 249 // this will throw error

	highScores = append(highScores, 555, 666, 777) // this won't
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
