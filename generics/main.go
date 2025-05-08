package main

import (
	"fmt"
)

func pritSlice[T int | string](items []T) {
	for _, value := range items {
		fmt.Println(value)
	}
}

func main() {
	pritSlice([]int{1, 2, 3, 4})
	pritSlice([]string{"abdullah", "Al", "Noman"})

}
