package main

import "fmt"

func changeSlice(s []int) []int {
	s = append(s, 100, 200, 300)
	return s
}

// variadic function
func variadicFunc(numbers ...int) {
	fmt.Println("numbers:", numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))
}

func main() {

	// empty slice

	var x []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	// p := changeSlice(x)
	// p := changeSlice(y)

	// fmt.Println("slice p:", p)

	fmt.Println("slice", x)
	fmt.Println("slice", y)

	// variadic function call
	variadicFunc(1, 2, 3, 4, 5)

	// fmt.Println("len", len(x))
	// fmt.Println("capacity", cap(x))
	// fmt.Println("slice", y)
	// fmt.Println("len", len(y))
	// fmt.Println("capacity", cap(y))

	// make function for slice
	// [1,2,3,4,5]
	// makeSlice := make([]int, 5, 8) // 5 is length , 8 is capacity
	// makeSlice[4] = 10

	// newSlice := makeSlice[3:5]

	// fmt.Println("slice ", newSlice)
	// fmt.Println("slice len", len(newSlice))
	// fmt.Println("slice cap", cap(newSlice))
	// fmt.Println("slice ", makeSlice)
	// fmt.Println("slice len", len(makeSlice))
	// fmt.Println("slice cap", cap(makeSlice))

	// s := []int{1, 2, 3, 4, 5}

	// fmt.Println("slice", s)
	// fmt.Println("len", len(s))
	// fmt.Println("capacity", cap(s))

	// after slice

	// sli := s[1:3]

	// fmt.Println("slice", sli)         // slice are the actual value
	// fmt.Println("len", len(sli))      // length is slich data length
	// fmt.Println("capacity", cap(sli)) //capacity is start time to end of the slice

}
