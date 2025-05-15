package main

import "fmt"

type Person struct {
	FristName string
	LastName  string
	Age       int
	Email     string
}

func main() {
	personOne := Person{
		FristName: "Abdullah",
		LastName:  "Al noman",
		Age:       28,
		Email:     "abdullahalnomancs@gmail.com",
	}
	fmt.Println(personOne)
}
