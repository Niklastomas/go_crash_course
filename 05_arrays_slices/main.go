package main

import "fmt"

func main() {
	//Arrays
	var fruits [2]string

	//Assign values
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	// Declare and assign
	cars := [2]string{"Audi", "Bmw"}

	fmt.Println(fruits)
	fmt.Println(cars[0])

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
}
