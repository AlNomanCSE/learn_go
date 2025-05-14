package main

import "fmt"

func createIDGenerator() func() int {
	nextId := 0
	return func() int {
		nextId++
		return nextId
	}
}

func main() {
	generateOrderID := createIDGenerator()
	fmt.Println(generateOrderID())
	fmt.Println(generateOrderID())
	fmt.Println(generateOrderID())
}
