package main

import (
	"fmt"
)

func main() {

	// Example 1
	// basic syntax
	// comprised of init, condition, and post statements
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// prints
	// 0
	// 1
	//  2
	//  3
	//  4

	// Example 2
	// initializing more than one variable in a for loop
	// go will read i++ or j++ as an individual statement and you can only have one statement per semi colon; hence  i, j = i+1, j+1
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4

	// Example 3
	// SCOPE
	// go is block scoped
	// in this example the variable "i" is scoped outside of the for block, whereas the previous examples are scooped ONLY to the for block.
	// syntax note: you need the ";" before the i<5 statement.
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("i can be used outside the for block", i)
	// 0
	// 1
	// 2
	// 3
	// 4
	// i can be used outside the for block 5

}
