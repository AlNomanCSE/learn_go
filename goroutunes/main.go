package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	// Start a goroutine
	go sayHello()

	// Main program prints something
	fmt.Println("Hello from main!")

	// Wait briefly to let goroutine finish
	time.Sleep(time.Millisecond * 100)
}
