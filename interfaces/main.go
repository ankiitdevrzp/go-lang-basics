package main

import (
	"fmt"
	"math"
)

// 2 classes circle and square

// define structs
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// define methods for square
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

// define methods for circle
func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// define a function that can accept both type of shapes
// else we have to create 2 functions
// so interface shape is created

type shape interface {
	area() float64
	circumference() float64
}

func getShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumference())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		square{length: 15.2},
		circle{radius: 5.2},
		circle{radius: 7.0},
	}

	for _, v := range shapes {
		getShapeInfo(v)
		fmt.Println("---")
	}
}
