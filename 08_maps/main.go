package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["Test1"] = "test1@test.com"
	emails["Test2"] = "test2@test.com"

	fmt.Println(emails)

	// Delete from map
	delete(emails, "Test1")
	fmt.Println(emails)

	// Pop value
	pop, ok := emails["Test2"]
	fmt.Println(pop, ok)

	// Declare map and add key values
	scores := map[string]int{"Bob": 10, "Joe": 20}
	fmt.Println(scores)
}
