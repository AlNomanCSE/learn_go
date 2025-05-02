package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange"}
	fruits = append(fruits, "grape", "pineapple")
	// var moreFruits []string = []string{"kiwi", "strawberry"}
	for in, fruit := range fruits {
		fmt.Printf("%d %s\n", in+1, fruit)
	}

	for i := range fruits {
		fmt.Println(fruits[i])
	}
}
