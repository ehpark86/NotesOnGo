package main

import "fmt"

/*
Notes:
 - Appending to an array does not create a new memory address.
 - Slicing an array create a NEW memory address. However, the sliced array
 has points to the original array that you sliced from. In other words, any modifications
 to the sliced array will affect the underlying original array.
 In essence, everytime you create a slice, you are copying the original array to a new memory address that points
 back to the underlying array.
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
