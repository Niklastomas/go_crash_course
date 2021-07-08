package main

import (
	"fmt"
	"sync"
)

// WaitGroup
var wg = sync.WaitGroup{}

func main() {
	// Goroutine
	msg := "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
}
