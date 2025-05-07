package main

import "fmt"

type Barker interface {
	Bark() string
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (c Cat) Bark() string {
	return fmt.Sprintf("%s says: Meow!", c.name)
}
func (d Dog) Bark() string {
	return fmt.Sprintf("%s says: Woof!", d.name)
}
func makeSound(b Barker) {
	fmt.Println(b.Bark())
}
func main() {
	dog := Dog{
		name: "Kutta",
	}
	cat := Cat{
		name: "billi",
	}
	makeSound(dog)
	makeSound(cat)

}
