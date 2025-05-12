package main

import "fmt"

var a int = 12

func main() {
	fmt.Printf("Hello from main %d", a)
}

func init() {
	a = a + 12
	fmt.Printf("From init %d", a)
}
