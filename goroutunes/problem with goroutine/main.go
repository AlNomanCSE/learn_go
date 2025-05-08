package main

import (
	"fmt"
	"sync"
)

func sayMessage(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	wg.Done() // Mark this goroutine as finished
}

func main() {
	var wg sync.WaitGroup // Create a Wait Group

	// Add 2 goroutines to the Wait Group
	wg.Add(2)

	// Start first goroutine
	go sayMessage("Hello from goroutine 1!", &wg)

	// Start second goroutine
	go sayMessage("Hello from goroutine 2!", &wg)

	// Wait for all goroutines to finish
	fmt.Println("Main: Waiting for goroutines...")
	wg.Wait()
	fmt.Println("Main: All goroutines done!")
}
