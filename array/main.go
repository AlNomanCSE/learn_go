package main

import "fmt"

func main() {
	var a [5]string = [5]string{"Hello", "Abdullah", "al", "Noman", "Boss"}
	for _, value := range a {
		fmt.Println(value)
	}
}
