package main

import (
	"fmt"
)

// custome package function
// import (
// 	"math-package/mathlib"
// )

// function for add

// func add(a int, b int) int {
// 	return a + b
// }

// func twoValueReturn(a int, b int) (int, int) {

// 	add := a + b
// 	mul := a * b
// 	return add, mul
// }

//! higher order function
//! function as a parameter
//! function as a return

// higher order function declare with function and paramerer

func higherOrderFunction(a, b int, Op func(y int, z int)) {
	Op(a, b)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

//! higher order function with return function

func higherOrderFunctionWithReturn(a, b int, Op func(y int, z int)) func(y int, z int) {
	Op(a, b)
	return add
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
	// add, mul := twoValueReturn(5, 10)
	// fmt.Println("return two value sum value and multiple of value", add, mul)

	// call package function
	// result := mathlib.Add(5, 10)

	// IIFE Immediately Invoked Function Expression
	// func() {
	// 	a := 10
	// 	b := 20
	// 	fmt.Println("a + b", a+b)
	// }()

	// fmt.Println("package function call result", result)

	//! higher order function call

	// higherOrderFunction(5, 10, add)

	//! higher order function with return
	// output will be 15 , 3
	res := higherOrderFunctionWithReturn(5, 10, add)

	res(1, 2)

}
