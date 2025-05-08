package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int // Shared variable
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect counter

	// Add 2 goroutines to WaitGroup
	wg.Add(2)

	// Goroutine 1: Increment counter 1000 times
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()   // Lock before accessing counter
			counter++   // Safely increment
			mu.Unlock() // Unlock after
		}
		wg.Done()
	}()

	// Goroutine 2: Increment counter 1000 times
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()   // Lock before accessing counter
			counter++   // Safely increment
			mu.Unlock() // Unlock after
		}
		wg.Done()
	}()

	// Wait for both goroutines to finish
	wg.Wait()
	fmt.Println("Final counter:", counter) // Should be 2000
}

// Small delay to ensure goroutine output is visible
//time.Sleep(time.Millisecond * 100)
