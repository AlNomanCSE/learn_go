package main

import "fmt"

type Operation func(int, int) int

func execute(op Operation, a, b int) int {
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

func main() {
	result := execute(add, 5, 3)
	fmt.Println(result)
}
