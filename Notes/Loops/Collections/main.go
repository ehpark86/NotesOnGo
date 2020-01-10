package main

import (
	"fmt"
)

func main() {
	// Example 1
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
		// 0 1
		// 1 2
		// 2 3
	}

	// if you only want the value then use the write only operator _
	// for _, v

	someString := "hello go"
	for k, v := range someString {
		fmt.Println(k, string(v))
		// 0 h
		// 1 e
		// 2 l
		// 3 l
		// 4 o
		// 5
		// 6 g
		// 7 o
	}
}
