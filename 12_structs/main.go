package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Gretting method (value reciever)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(lastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = lastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Samantha", lastName: "Smith", city: "Boston",
		gender: "f", age: 27,
	}
	// Alternative
	person2 := Person{"Niklas", "Tomas", "Stenstorp", "m", 27}

	person1.getMarried("Tomas")
	person2.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.age)
}
