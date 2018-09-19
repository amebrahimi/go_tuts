package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "Mike@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@mgail.com", "Sharon": "Sharon@gmail.com"}

	emails["Mike"] = "Mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
}
