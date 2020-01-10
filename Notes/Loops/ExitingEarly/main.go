package main

import (
	"fmt"
)

func main() {
	// Go's implementation of Do While loop
	i := 0
	for i < 5 {
		i++
		if i == 5 {
			break
		}
		fmt.Println(i)
		// 1
		// 2
		// 3
		// 4
	}

	// CONTINUE
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			// if i is evenly divisible by 2 then go back to the loop.
			// exit this itertion of the loop
			continue
		}
		fmt.Println(i)
		// 1
		// 3
	}

	// LABLES to break out of outer for loop
Loop:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i * j)
			// 1
			// 2
			// 3
			if i*j >= 3 {
				// break out of the loop with the label Loop
				break Loop
			}
		}
	}

}
