package main

import "fmt"

func main() {
	printInteger := func(d int) int {
		return d
	}
	fmt.Println(printInteger(12))
	myStrgin1, myStrinn2 := swapingString("noman", "abdullah")
	fmt.Println(myStrgin1, myStrinn2)
	fn := func(a int) int {
		return a * 5
	}
	functionTakingFunc(fn)
	variadictFunc(12, 3, 4, 5, 5, 5)
}

func swapingString(s1 string, s2 string) (string, string) {
	return s2, s1
}

func functionTakingFunc(fn func(a int) int) {
	fmt.Println(fn(5))
}

func variadictFunc(nums ...int) {
	for _, value := range nums {
		fmt.Println(value + 10)
	}
}
