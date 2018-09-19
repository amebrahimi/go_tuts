package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int

	// Using var
	var age int32 = 32
	const isCool = true
	var size float32 = 2.3

	name, email := "Brad", "amirebrahimi5@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", isCool)
}
