package main

import "fmt"

/*
A map takes a key : value.
Map one key type and one value type.
When declaring a map, the type of the key and value must be consistent with every element.
The return order is not guaranteed.

adding
mapName[key] = value

deleting
delete(nameOfMap, key)

checking if values exits
_, ok := mapName[keyName] // if the key doesn't exist, ok will print out false
*/

func main() {
	// the literal syntax
	// each key must be a string and each value that maps to it must be an int
	// a trailing comma is needed
	ageOfPeople := map[string]int{
		"Eugene": 34,
		"Sam":    65,
		"Linda":  63,
		"John":   32,
	}
	// m := map[[]int]string{}  will throw an error because a slice cannot be a key to a map.
	// m := map[[3]int]string{} this will work because the key is now an array.
	fmt.Println(ageOfPeople)

}
