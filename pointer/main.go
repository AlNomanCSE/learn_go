package main

import "fmt"

func tryChnage(value *int) {
	*value = 100
}
func main() {
	x := 10
	p := &x
	tryChnage(&x)
	fmt.Println(x, *p)
}
