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

	// Declare map and add key values
	scores := map[string]int{"Bob": 10, "Joe": 20}
	print(scores)

}
