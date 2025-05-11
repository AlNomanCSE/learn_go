package main

import (
	"fmt"
)

type greeter struct{}

func (g greeter) sayHello() string {
	return "Hello from a method!"
}

func main() {
	g := greeter{}
	fmt.Println(g.sayHello())
	// mathlib.Add(12, 13)
	add()
}
