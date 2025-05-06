package main

import "fmt"

func main() {
	myMap1 := map[string]int{
		"price": 12300,
		"name":  111,
	}
	myMap2 := make(map[string]int)
	fmt.Println(myMap1, myMap2)
	for key, value := range myMap1 {
		myMap2[key] = value
	}
	fmt.Println(myMap1, myMap2)
}
