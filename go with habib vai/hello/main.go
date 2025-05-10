package main

import "fmt"

type greeter struct{}

func (g greeter) sayHello() {
	fmt.Println("Hello from a method!")
}
func main() {
	age := 18
	fmt.Println(age)
	g := greeter{}
	g.sayHello()
}
