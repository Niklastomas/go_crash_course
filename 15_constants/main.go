package main

import "fmt"

// Value must be calculated at comple time
// PascalCase for exported constants
// camelCase for internal constants
// Can interoperate with similar types

// constant
const pi = 3.14

// enum
const (
	North = iota
	South
	East
	West
)

type Direction int

func main() {
	// fmt.Println(pi)

	var myDirection Direction
	myDirection = North

	switch myDirection {
	case North:
		fmt.Println("myDirection is North: ", myDirection)
	case South:
		fmt.Println("myDirection is South: ", myDirection)
	case East:
		fmt.Println("myDirection is East: ", myDirection)
	case West:
		fmt.Println("myDirection is West: ", myDirection)
	default:
		fmt.Println("myDirection is Unknown: ", myDirection)

	}

}
