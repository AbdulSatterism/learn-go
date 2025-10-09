package main

import (
	"fmt"
)

// function for add

func add(a int, b int) int {
	return a + b
}

func twoValueReturn(a int, b int) (int, int) {

	add := a + b
	mul := a * b
	return add, mul
}

func main() {
	// fmt.Println("hello go-lang this is my new journy")

	// a := 10

	// a = 10.5  // error can't used float in int
	// a = 20 // ok

	// fmt.Println("my value is ", a)

	// call add funtinon
	// var result = add(5, 10)

	// fmt.Println(result)

	// call add function with short hand
	// result := add(5, 10)
	// fmt.Println(result)

	// call second value return function
	add, mul := twoValueReturn(5, 10)
	fmt.Println("return two value sum value and multiple of value", add, mul)
}
