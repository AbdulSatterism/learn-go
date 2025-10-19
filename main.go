package main

import "fmt"

// learn pointer in go-lang

// call by value

func callByValue(num int) {
	fmt.Println("value inside function callByValue", num)
}

// call by reference

func callByReference(ptr *int) {
	fmt.Println("address inside function callByReference", ptr)
	fmt.Println("value inside function callByReference", *ptr)
}

// call by reference different way

// func refArray(arr [5]int) {

// 	fmt.Println("array inside function callByReference", arr)

// }

func refArray(arr *[5]int) {

	fmt.Println("array inside function callByReference", *arr)

}

func main() {

	// num := 50

	// ptr := &num

	// fmt.Println("address of num before pointer", ptr) // address of num

	// fmt.Println("value of num before pointer", *ptr) // dereference value

	// callByValue(num)

	// callByValue(num)

	// call by reference

	// callByReference(&num)

	// call by reference different way

	arr := [5]int{1, 2, 3, 4, 5}

	refArray(&arr)

}

/*
// learn struct and receiver function

type TUser struct {
	Name string
	Age  int
}

// normal function
func personDetails(user TUser) {
	fmt.Println("Name: ", user.Name, "Age: ", user.Age)
}

// receiver function

func (user TUser) person() {
	fmt.Println("Name: ", user.Name, "Age: ", user.Age)
}

func main() {

	user1 := TUser{
		Name: "Mr. Kodu",
		Age:  30,
	}

	// call receiver function
	user1.person()

	// personDetails(user1)

	// user2 := TUser{
	// 	Name: "Mr. Jodu",
	// 	Age:  25,
	// }

	// personDetails(user2)

	// fmt.Println("Name: ", user1.Name, "Age: ", user1.Age)
	// fmt.Println("Name: ", user2.Name, "Age: ", user2.Age)

}

/*

// closure learn program

const a = 100

var p = 10

func outer() func() {
	money := 100

	age := 18

	fmt.Println("I am over ", age)

	show := func() {
		money = money + a + p

		fmt.Println(money)
	}

	return show
}

func call() {
	show1 := outer()

	show1()
	show1()

	show2 := outer()
	show2()
}

func main() {
	call()
}

func init() {
	fmt.Println("===> Bank Account <===")
}

/*

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


*/
