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
	fmt.Println(cars[0])
	numbers := [...]int{1, 2, 3, 4}
	fmt.Println(numbers)

	// Arrays fo arrays
	identityMatrix := [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	//Slices
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6t elements
	e := a[3:6] //slice hte 4th, 5th and 6th elements
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Built in functions
	fruitSlice := make([]string, 3, 50)
	fruitSlice = append(fruitSlice, "Apple", "Orange")
	capacity := cap(fruitSlice)
	numOfFruits := len(cars)
	fmt.Println(fruitSlice)
	fmt.Printf("Numbers of fruits: %d\n", numOfFruits)
	fmt.Printf("Capacity: %d\n", capacity)

}
