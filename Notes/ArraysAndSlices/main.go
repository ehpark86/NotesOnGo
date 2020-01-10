package main

import "fmt"

/*
Notes:
 - Appending to an array does not create a new memory address.
 - Slicing an array creates a NEW memory address. However, the sliced array
	points to the original array that was sliced from. In other words, any modifications
	to the sliced array will affect the underlying original array.
 	In essence, everytime you create a slice, you are copying the original array to a new memory address that points
	 back to the underlying array.

	3 different ways to delcare an array
	===========================
	a := [3]int{1,2,3}
	a := [...]int{1,2,3} // if you need to be flexible with the size of the array.
	var a [3]int // each index is set to 0 by default

	len() function returns size of array
	copies refer to the different underyling data -- if you have a 3 element arrray and you assign that to another variable,
	all 3 elements are going to be copied to a new memory location and you're working with an independent copy. Could be very expensive
	e.g.
	a and b are stored in a different memory address.
	what you do to b won't change a.
	a := [3]int{1, 2, 3}
	b := a

Slices
- Backed by an array. Every slice you create Go has an array created behind the scenes.
- make(type of object, length, capacity) function

		a := make([]int, 10) // create slice of type int with capacity and length == 10
		a := make([]int, 10, 100) // length == 10, capacity == 100

*/
func main() {
	// declare and initialize an empty array
	a := []int{}
	fmt.Printf("%p\n", &a) // 0xc00009a020

	// append 1 and 2 to the empty array  "a"
	a = append(a, 1, 2)
	fmt.Printf("%p\n", &a) // 0xc00009a020

	b := a[1:]
	b[0] = 5
	fmt.Println(b)         // [5]
	fmt.Println(a)         // [1,5]
	fmt.Printf("%p\n", &b) // 0xc00009a040

	c := b[1:]
	fmt.Printf("%p\n", &c) // 0xc00009a0a0
}
