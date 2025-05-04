package main

import (
	"fmt"
)

func main() {
	const (
		core = 500
		host = "Localhost"
	)
	var (
		isNameAvailable = false
		isString        = "yes"
	)
	fruits := deck{"apple", "banana", "orange"}
	fruits = append(fruits, "grape", "pineapple")

	hand, remainingCrads := deal(fruits, 3)
	fmt.Println(hand, remainingCrads)
	fmt.Println(isNameAvailable, isString)
	fmt.Println(core, host)
	forAsWhile()
	classicForLoop()
}

func forAsWhile() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
}

func classicForLoop() {
	var (
		even   []int
		divBy3 []int
	)

	for j := range 10 {
		if j%2 == 0 {
			even = append(even, j)
		} else if j%3 == 0 {
			divBy3 = append(divBy3, j)
		} else {
			fmt.Println("Print nothing")
		}
	}
}
