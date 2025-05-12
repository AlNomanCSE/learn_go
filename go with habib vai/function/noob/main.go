package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	fmt.Printf("%d \n", add(12, 14))
}
