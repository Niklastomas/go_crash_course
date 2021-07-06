package main

import "fmt"

func main() {
	// Slice
	ids := []int{33, 55, 23, 63, 89}

	// Loop trough ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	scores := map[string]int{"Bob": 10, "Joe": 20}

	for key, value := range scores {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

}
