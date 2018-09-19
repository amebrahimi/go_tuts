package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// getsMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	//Init person using strict
	// person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 12}

	// Alternative
	person1 := Person{"Samant", "smith", "Bost", "f", 12}
	fmt.Println(person1)

	person1.hasBirthday()

	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
}
