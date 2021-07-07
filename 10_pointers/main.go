package main

import "fmt"

func main() {
	a := 5
	// use & to get address of variable
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address(Dereference pointer)
	fmt.Println(*b)

	// Change value with pointer
	*b = 10
	fmt.Println(a)

	// Complex types (e.g. structs) are automatically dereferenced
	var car *Car
	// car = &Car{brand: "Volvo", color: "Red"}
	car = new(Car)
	car.brand = "Volvo"
	fmt.Println(car.brand)

}

type Car struct {
	brand string
	color string
}
