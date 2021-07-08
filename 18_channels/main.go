package main

import (
	"fmt"
	"sync"
)

// Create channel - make(chan int)
// Send message into channel - ch <- val
// Receive message from channel - val := <- ch
// Send only: chan <- int
// Receive only: <-chan int
// Create buffered channel - make(chan int, 50)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	// Receiving goroutine
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		// i := <-ch

		wg.Done()

	}(ch)

	// Sending goroutine
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

}
