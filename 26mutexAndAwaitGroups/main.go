package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - LearnCodeOnline.in")

	wg := &sync.WaitGroup{}

	var score = []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		score = append(score, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		score = append(score, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
