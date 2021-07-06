package main

import "fmt"

// MAIN TYPES
// string
// bool
// int
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

func main() {
	var name string = "Niklas"
	var age int = 27
	var isHuman = true
	firstname, lastname := "Niklas", "Tomas"

	fmt.Println(name, firstname, lastname, age, isHuman)
	fmt.Printf("Name: %T\nAge :%T\n", name, age)
}
