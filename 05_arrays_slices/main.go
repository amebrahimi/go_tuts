package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArray [2]string

	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// fruitArray := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	//Range
	fmt.Println(fruitSlice[1:2])
}
