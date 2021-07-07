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

	peopleScore := map[string]int{"Joe": 20, "Doe": 50, "Sam": 77}
	// if with initializer
	if score, ok := peopleScore["Sam"]; ok {
		fmt.Println(score)
	}

	// Switch
	color := "white"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	case "black", "white":
		fmt.Println("Color is black or white")
	default:
		fmt.Println("Color is not red or blue")

	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is and int")
	case float64:
		fmt.Println("i isn an float64")
	case string:
		fmt.Println("i is and string")
	default:
		fmt.Println("i is another type")

	}

	message := "Hej"
	fmt.Printf("%v, %T\n", message[0], string(message[0]))

}
