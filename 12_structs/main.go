package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	firstName string
	lastName  string
}

type Person struct {
	Human
	city   string
	gender string
	age    int
}

// struct with tag
type Book struct {
	name  string `required max:"100"`
	pages int
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
		Human:  Human{firstName: "Samantha", lastName: "Smith"},
		gender: "f", age: 27, city: "Boston",
	}
	// Alternative
	person2 := Person{
		Human{"Niklas", "Tomas"},
		"Stenstorp", "m", 27,
	}

	person1.getMarried("Tomas")
	person2.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.age)
}
