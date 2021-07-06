package main

import "fmt"

func main() {

	x := 15
	y := 10

	// if else
	if x < 10 {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)

	}

	// Switch
	color := "blue"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")

	}

	message := "Hej"
	fmt.Printf("%v, %T\n", message[0], string(message[0]))

}
