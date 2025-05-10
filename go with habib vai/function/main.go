package main

import "fmt"

type greeter struct{}

func (g greeter) sayHello() string {
	return "Hello from a method!"
}
func (g greeter) sayHello() int {
	return 123
}
func main() {
	g := greeter{}
	fmt.Println(g.sayHello())

}
